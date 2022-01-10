package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/mongo"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
)

func NewSaveProfileUseCase(logger *common.Logger, cacheGateway *redis.Gateway, databaseGateway *mongo.Gateway, getAllNonFungibleTokens GetAllNonFungibleTokensUseCase) SaveProfileUseCase {
	return SaveProfile(logger, cacheGateway, databaseGateway, getAllNonFungibleTokens)
}

// fetch metadata and ownership of nfts being submitted
// try to save the profile to the cache
// regardless of error, save the profile to the database
func SaveProfile(logger *common.Logger, cacheGateway gateways.ICacheGateway, databaseGateway gateways.IDatabaseGateway, getAllNonFungibleTokens GetAllNonFungibleTokensUseCase) SaveProfileUseCase {
	return func(ctx context.Context, profile *entities.Profile) error {
		if err := common.ValidateStruct(profile); err != nil {
			return err
		}

		profile.NonFungibleTokens = getAllNonFungibleTokens(ctx, profile.Address, profile.NonFungibleTokens)

		cacheGateway.SaveProfile(ctx, profile)

		return databaseGateway.SaveProfile(ctx, profile)
	}
}
