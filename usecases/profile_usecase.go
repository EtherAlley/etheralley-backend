package usecases

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/eflem00/go-example-app/common"
	"github.com/eflem00/go-example-app/entities"
	"github.com/eflem00/go-example-app/gateways/cache"
	"github.com/eflem00/go-example-app/gateways/db"
)

type ProfileUsecase struct {
	logger            *common.Logger
	cache             *cache.Cache
	profileRepository *db.ProfileRepository
}

func NewProfileUseCase(logger *common.Logger, cache *cache.Cache, profileRepository *db.ProfileRepository) *ProfileUsecase {
	return &ProfileUsecase{
		logger,
		cache,
		profileRepository,
	}
}

// check cache for key and touch if we get a cache hit
// if cache miss, go to persistant storage and set
func (uc *ProfileUsecase) GetProfileByAddress(ctx context.Context, address string) (entities.Profile, error) {
	profileString, err := uc.cache.Get(ctx, address)

	// should check the type of error for redis.Nil here but we'll keep it simple and treat this as a cache miss
	if err != nil {
		uc.logger.Debugf("Cache miss for key %v", address)

		profile, err := uc.profileRepository.GetProfileByAddress(address)

		if err != nil {
			return entities.Profile{}, errors.New("no value for provided key")
		}

		uc.cache.Set(ctx, address, profile, time.Hour)

		return profile, nil
	} else { // cache hit, use the value and touch the key
		uc.logger.Debugf("Cache hit for key %v", address)

		uc.cache.Touch(ctx, address)

		profile := entities.Profile{}
		json.Unmarshal([]byte(profileString), &profile)

		return profile, nil
	}
}

func (uc *ProfileUsecase) SaveProfile(ctx context.Context, address string, profile entities.Profile) error {
	profileBytes, err := json.Marshal(profile)

	if err == nil {
		uc.cache.Set(ctx, address, string(profileBytes), time.Hour)
	}

	return uc.profileRepository.SaveProfile(address, profile)
}
