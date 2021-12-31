package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/mongo"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
)

func NewSaveProfileUseCase(logger *common.Logger, cacheGateway *redis.Gateway, databaseGateway *mongo.Gateway, getNFTUseCase GetNFTUseCase) SaveProfileUseCase {
	return SaveProfile(logger, cacheGateway, databaseGateway, getNFTUseCase)
}

// TODO: fetch metadata and ownership of nfts being submitted (concurrently)
// try to save the profile to the cache
// regardless of error, save the profile to the database
func SaveProfile(logger *common.Logger, cacheGateway gateways.ICacheGateway, databaseGateway gateways.IDatabaseGateway, getNFTUseCase GetNFTUseCase) SaveProfileUseCase {
	return func(ctx context.Context, profile *entities.Profile) error {
		nfts := []entities.NFT{}
		for _, partialNft := range profile.NFTs {
			nft, err := getNFTUseCase(ctx, profile.Address, partialNft.Blockchain, partialNft.ContractAddress, partialNft.SchemaName, partialNft.TokenId)
			if err == nil && nft.Owned {
				nfts = append(nfts, *nft)
			} else {
				logger.Debugf("invalid nft provided: %v", err)
			}
		}

		profile.NFTs = nfts

		cacheGateway.SaveProfile(ctx, profile)

		return databaseGateway.SaveProfile(ctx, profile)
	}
}
