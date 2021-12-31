package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/mongo"
	"github.com/etheralley/etheralley-core-api/gateways/opensea"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
)

func NewGetProfileUsecase(logger *common.Logger, cacheGateway *redis.Gateway, databaseGateway *mongo.Gateway, nftMarketGateway *opensea.Gateway) GetProfileUsecase {
	return GetProfile(logger, cacheGateway, databaseGateway, nftMarketGateway)
}

// first try to get the profile from the cache.
// if cache miss, go to database
// if database miss, build default from open sea
func GetProfile(logger *common.Logger, cacheGateway gateways.ICacheGateway, databaseGateway gateways.IDatabaseGateway, nftMarketGateway gateways.INFTMarketGateway) GetProfileUsecase {
	return func(ctx context.Context, address string) (*entities.Profile, error) {
		profile, err := cacheGateway.GetProfileByAddress(ctx, address)

		if err == nil {
			logger.Debugf("cache hit for address %v, returning", address)
			return profile, nil
		}

		logger.Debugf("cache miss for address %v, going to db", address)

		profile, err = databaseGateway.GetProfileByAddress(ctx, address)

		if err == common.ErrNil {
			logger.Debugf("db miss for address %v, building default", address)
			nfts, err := nftMarketGateway.GetNFTs(address)

			if err != nil {
				logger.Err(err, "err calling nft market gateway")
				return nil, err
			}

			profile = &entities.Profile{
				Address: address,
				NFTs:    nfts,
			}

			cacheGateway.SaveProfile(ctx, profile)
			return profile, nil
		}

		// TODO: validate the nfts in the profile are still owned
		logger.Debugf("db hit for address %v, validating nft ownership", address)
		cacheGateway.SaveProfile(ctx, profile)

		return profile, nil
	}
}
