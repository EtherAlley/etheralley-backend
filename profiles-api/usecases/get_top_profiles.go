package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
	"github.com/etheralley/etheralley-backend/profiles-api/gateways"
)

func NewGetTopProfilesUseCase(
	logger common.ILogger,
	cacheGateway gateways.ICacheGateway,
	getLightProfile IGetLightProfileUseCase,
) IGetTopProfilesUseCase {
	return &getTopProfilesUseCase{
		logger,
		cacheGateway,
		getLightProfile,
	}
}

type getTopProfilesUseCase struct {
	logger          common.ILogger
	cacheGateway    gateways.ICacheGateway
	getLightProfile IGetLightProfileUseCase
}

type IGetTopProfilesUseCase interface {
	Do(ctx context.Context, input *GetTopProfilesInput) *[]entities.Profile
}

type GetTopProfilesInput struct {
}

func (uc *getTopProfilesUseCase) Do(ctx context.Context, input *GetTopProfilesInput) *[]entities.Profile {
	if profiles, err := uc.cacheGateway.GetProfiles(ctx, "top"); err == nil {
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

			profile, err := uc.getLightProfile.Do(ctx, &GetLightProfileInput{
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

	uc.cacheGateway.SaveProfiles(ctx, "top", &trimmedProfiles)

	return &trimmedProfiles
}
