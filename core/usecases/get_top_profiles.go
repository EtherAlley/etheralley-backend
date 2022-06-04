package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/core/entities"
	"github.com/etheralley/etheralley-core-api/core/gateways"
)

func NewGetTopProfilesUseCase(
	logger common.ILogger,
	cacheGateway gateways.ICacheGateway,
	getProfile IGetProfileUseCase,
) IGetTopProfilesUseCase {
	return &getTopProfilesUseCase{
		logger,
		cacheGateway,
		getProfile,
	}
}

type getTopProfilesUseCase struct {
	logger       common.ILogger
	cacheGateway gateways.ICacheGateway
	getProfile   IGetProfileUseCase
}

type IGetTopProfilesUseCase interface {
	Do(ctx context.Context, input *GetTopProfilesInput) *[]entities.Profile
}

type GetTopProfilesInput struct {
}

func (uc *getTopProfilesUseCase) Do(ctx context.Context, input *GetTopProfilesInput) *[]entities.Profile {
	if profiles, err := uc.cacheGateway.GetTopViewedProfiles(ctx); err == nil {
		return profiles
	}

	addresses, err := uc.cacheGateway.GetTopViewedAddresses(ctx)

	if err != nil {
		return &[]entities.Profile{}
	}

	var wg sync.WaitGroup
	profiles := make([]*entities.Profile, len(*addresses))

	for i, address := range *addresses {
		wg.Add(1)

		go func(i int, address string) {
			defer wg.Done()

			profile, err := uc.getProfile.Do(ctx, &GetProfileInput{
				Address: address,
			})

			if err != nil {
				uc.logger.Warn(ctx).Err(err).Msgf("err hydrating top profile %v", address)
				return
			}

			if profile.Banned {
				uc.logger.Info(ctx).Msgf("excluding banned profile %v", address)
				return
			}

			profiles[i] = profile
		}(i, address)
	}

	wg.Wait()

	// trim any profiles that had an error fetching
	trimmedProfiles := []entities.Profile{}
	for _, profile := range profiles {
		if profile != nil {
			trimmedProfiles = append(trimmedProfiles, *profile)
		}
	}

	uc.cacheGateway.SaveTopViewedProfiles(ctx, &trimmedProfiles)

	return &trimmedProfiles
}
