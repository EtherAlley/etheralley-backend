package usecases

import (
	"context"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/entities"
	"github.com/eflem00/go-example-app/gateways/mongo"
	"github.com/eflem00/go-example-app/gateways/redis"
)

type ProfileUsecase struct {
	logger            *common.Logger
	profileCache      *redis.ProfileCache
	profileRepository *mongo.ProfileRepository
}

func NewProfileUseCase(logger *common.Logger, profileCache *redis.ProfileCache, profileRepository *mongo.ProfileRepository) *ProfileUsecase {
	return &ProfileUsecase{
		logger,
		profileCache,
		profileRepository,
	}
}

// check cache for key and touch if we get a cache hit
// if cache miss, go to persistant storage and set
func (uc *ProfileUsecase) GetProfileByAddress(ctx context.Context, address string) (entities.Profile, error) {
	profile, err := uc.profileCache.GetProfileByAddress(ctx, address)

	// should check the type of error for redis.Nil here but we'll keep it simple and treat this as a cache miss
	if err != nil {
		uc.logger.Debugf("Cache miss for address %v", address)

		profile, err := uc.profileRepository.GetProfileByAddress(ctx, address)

		if err != nil {
			return entities.Profile{}, err
		}

		uc.profileCache.SaveProfile(ctx, profile)

		return profile, nil
	}

	// cache hit, use the value and touch the key
	uc.logger.Debugf("Cache hit for address %v", address)

	return profile, nil
}

func (uc *ProfileUsecase) SaveProfile(ctx context.Context, profile entities.Profile) error {
	err := uc.profileCache.SaveProfile(ctx, profile)

	if err != nil {
		uc.logger.Debugf("Cache error for address %v: %v", profile.Address, err)
	}

	return uc.profileRepository.SaveProfile(ctx, profile)
}
