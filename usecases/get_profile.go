package usecases

import (
	"context"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/entities"
	"github.com/eflem00/go-example-app/gateways"
	"github.com/eflem00/go-example-app/gateways/mongo"
	"github.com/eflem00/go-example-app/gateways/redis"
)

type IGetProfileUsecase interface {
	Go(ctx context.Context, address string) (*entities.Profile, error)
}

type GetProfileUsecase struct {
	logger          *common.Logger
	cacheGateway    gateways.ICacheGateway
	databaseGateway gateways.IDatabaseGateway
}

func NewGetProfileUsecase(logger *common.Logger, cacheGateway *redis.Gateway, databaseGateway *mongo.Gateway) *GetProfileUsecase {
	return &GetProfileUsecase{
		logger,
		cacheGateway,
		databaseGateway,
	}
}

// check cache for key
// if cache miss, go to db and set in cache
func (uc *GetProfileUsecase) Go(ctx context.Context, address string) (*entities.Profile, error) {
	profile, err := uc.cacheGateway.GetProfileByAddress(ctx, address)

	if err == nil {
		uc.logger.Debugf("Cache hit for address %v", address)
		return profile, nil
	}

	uc.logger.Debugf("Cache miss for address %v", address)

	profile, err = uc.databaseGateway.GetProfileByAddress(ctx, address)

	if err != nil {
		return profile, err
	}

	uc.cacheGateway.SaveProfile(ctx, profile)

	return profile, nil
}
