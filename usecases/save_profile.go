package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/mongo"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
)

func NewSaveProfileUseCase(logger *common.Logger, cacheGateway *redis.Gateway, databaseGateway *mongo.Gateway, getAllNFTs GetAllNFTsUseCase) SaveProfileUseCase {
	return SaveProfile(logger, cacheGateway, databaseGateway, getAllNFTs)
}

// fetch metadata and ownership of nfts being submitted
// try to save the profile to the cache
// regardless of error, save the profile to the database
func SaveProfile(logger *common.Logger, cacheGateway gateways.ICacheGateway, databaseGateway gateways.IDatabaseGateway, getAllNFTs GetAllNFTsUseCase) SaveProfileUseCase {
	return func(ctx context.Context, profile *entities.Profile) error {
		nftLocations := &[]entities.NFTLocation{}
		for _, nft := range *profile.NFTs {
			*nftLocations = append(*nftLocations, *nft.Location)
		}

		profile.NFTs = getAllNFTs(ctx, profile.Address, nftLocations)

		cacheGateway.SaveProfile(ctx, profile)

		return databaseGateway.SaveProfile(ctx, profile)
	}
}
