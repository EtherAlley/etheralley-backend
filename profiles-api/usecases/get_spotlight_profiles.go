package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
	"github.com/etheralley/etheralley-backend/profiles-api/gateways"
	"github.com/etheralley/etheralley-backend/profiles-api/settings"
)

func NewGetSpotlightProfilesUseCase(
	logger common.ILogger,
	settings settings.ISettings,
	getLightProfile IGetLightProfileUseCase,
	cacheGateway gateways.ICacheGateway,
) IGetSpotlightProfilesUseCase {
	return &getSpotlightProfilesUseCase{
		logger,
		settings,
		getLightProfile,
		cacheGateway,
	}
}

type getSpotlightProfilesUseCase struct {
	logger          common.ILogger
	settings        settings.ISettings
	getLightProfile IGetLightProfileUseCase
	cacheGateway    gateways.ICacheGateway
}

type IGetSpotlightProfilesUseCase interface {
	Do(ctx context.Context, input *GetSpotlightProfilesInput) *[]entities.Profile
}

type GetSpotlightProfilesInput struct {
}

// Read the spotlight profile addresses from the env variables and call the standard get profile usecase to get a fully hydrated profile
func (uc *getSpotlightProfilesUseCase) Do(ctx context.Context, input *GetSpotlightProfilesInput) *[]entities.Profile {
	if profiles, err := uc.cacheGateway.GetProfiles(ctx, "spotlight"); err == nil {
		return profiles
	}

	addresses := uc.settings.SpotlightProfileAddresses()

	var wg sync.WaitGroup
	profiles := make([]*entities.Profile, len(addresses))

	for i, address := range addresses {
		wg.Add(1)

		go func(i int, address string) {
			defer wg.Done()

			profile, err := uc.getLightProfile.Do(ctx, &GetLightProfileInput{
				Address: address,
			})

			if err != nil {
				uc.logger.Warn(ctx).Err(err).Msgf("err hydrating spotlight profiles %v", address)
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

	uc.cacheGateway.SaveProfiles(ctx, "spotlight", &trimmedProfiles)

	return &trimmedProfiles
}
