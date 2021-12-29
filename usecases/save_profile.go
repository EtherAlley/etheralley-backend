package usecases

import (
	"context"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/entities"
	"github.com/eflem00/go-example-app/gateways"
	"github.com/eflem00/go-example-app/gateways/mongo"
	"github.com/eflem00/go-example-app/gateways/redis"
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
