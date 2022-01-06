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

func NewGetProfileUseCase(logger *common.Logger, cacheGateway *redis.Gateway, databaseGateway *mongo.Gateway, nftApiGateway *opensea.Gateway, getAllNFTs GetAllNFTsUseCase) GetProfileUseCase {
	return GetProfile(logger, cacheGateway, databaseGateway, nftApiGateway, getAllNFTs)
}

// first try to get the profile from the cache.
// if cache miss, go to database
// if database miss, build default
// if database hit, re-check ownership
func GetProfile(logger *common.Logger, cacheGateway gateways.ICacheGateway, databaseGateway gateways.IDatabaseGateway, nftApiGateway gateways.INFTAPIGateway, getAllNFTs GetAllNFTsUseCase) GetProfileUseCase {
	return func(ctx context.Context, address string) (*entities.Profile, error) {
		profile, err := cacheGateway.GetProfileByAddress(ctx, address)

		if err == nil {
			logger.Debugf("cache hit for profile %v", address)
			return profile, nil
		}

		logger.Debugf("cache miss for profile %v", address)

		profile, err = databaseGateway.GetProfileByAddress(ctx, address)

		if err == common.ErrNil {
			logger.Debugf("db miss for profile %v", address)
			nfts, err := nftApiGateway.GetNFTs(address)

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

		if err != nil {
			logger.Err(err, "err getting profile from db")
			return nil, err
		}

		logger.Debugf("db hit for profile %v", address)

		nftLocations := &[]entities.NFTLocation{}
		for _, nft := range *profile.NFTs {
			*nftLocations = append(*nftLocations, *nft.Location)
		}

		profile.NFTs = getAllNFTs(ctx, profile.Address, nftLocations)

		cacheGateway.SaveProfile(ctx, profile)

		return profile, nil
	}
}
