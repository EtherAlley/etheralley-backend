package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/mongo"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
)

func NewSaveProfileUseCase(logger *common.Logger, cacheGateway *redis.Gateway, databaseGateway *mongo.Gateway, hydrateNFTs HydrateNFTsUseCase) SaveProfileUseCase {
	return SaveProfile(logger, cacheGateway, databaseGateway, hydrateNFTs)
}

// TODO: fetch metadata and ownership of nfts being submitted (concurrently)
// try to save the profile to the cache
// regardless of error, save the profile to the database
func SaveProfile(logger *common.Logger, cacheGateway gateways.ICacheGateway, databaseGateway gateways.IDatabaseGateway, hydrateNFTs HydrateNFTsUseCase) SaveProfileUseCase {
	return func(ctx context.Context, profile *entities.Profile) error {
		profile.NFTs = hydrateNFTs(ctx, profile.Address, profile.NFTs)

		cacheGateway.SaveProfile(ctx, profile)

		return databaseGateway.SaveProfile(ctx, profile)
	}
}
