package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/mongo"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
)

func NewSaveProfileUseCase(logger *common.Logger, cacheGateway *redis.Gateway, databaseGateway *mongo.Gateway) SaveProfileUseCase {
	return SaveProfile(logger, cacheGateway, databaseGateway)
}

// try to save the profile to the cache
// regardless of error, save the profile to the database
func SaveProfile(logger *common.Logger, cacheGateway gateways.ICacheGateway, databaseGateway gateways.IDatabaseGateway) SaveProfileUseCase {
	return func(ctx context.Context, profile *entities.Profile) error {
		err := cacheGateway.SaveProfile(ctx, profile)

		if err != nil {
			logger.Debugf("Cache error for address %v: %v", profile.Address, err)
		}

		return databaseGateway.SaveProfile(ctx, profile)
	}
}
