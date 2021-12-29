package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/mongo"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
)

type ISaveProfileUseCase interface {
	Go(ctx context.Context, profile *entities.Profile) error
}

type SaveProfileUseCase struct {
	logger          *common.Logger
	cacheGateway    gateways.ICacheGateway
	databaseGateway gateways.IDatabaseGateway
}

func NewSaveProfileUseCase(logger *common.Logger, cacheGateway *redis.Gateway, databaseGateway *mongo.Gateway) *SaveProfileUseCase {
	return &SaveProfileUseCase{
		logger,
		cacheGateway,
		databaseGateway,
	}
}

func (uc *SaveProfileUseCase) Go(ctx context.Context, profile *entities.Profile) error {
	err := uc.cacheGateway.SaveProfile(ctx, profile)

	if err != nil {
		uc.logger.Debugf("Cache error for address %v: %v", profile.Address, err)
	}

	return uc.databaseGateway.SaveProfile(ctx, profile)
}
