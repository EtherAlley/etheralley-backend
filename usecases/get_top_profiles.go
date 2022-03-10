package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type GetTopProfilesInput struct {
}

type IGetTopProfilesUseCase func(ctx context.Context, input *GetTopProfilesInput) *[]entities.Profile

func NewGetTopProfilesUseCase(
	logger common.ILogger,
	cacheGateway gateways.ICacheGateway,
	getProfile IGetProfileUseCase,
) IGetTopProfilesUseCase {
	return func(ctx context.Context, input *GetTopProfilesInput) *[]entities.Profile {
		if profiles, err := cacheGateway.GetTopViewedProfiles(ctx); err == nil {
			return profiles
		}

		addresses, err := cacheGateway.GetTopViewedAddresses(ctx)

		if err != nil {
			return &[]entities.Profile{}
		}

		var wg sync.WaitGroup
		profiles := make([]*entities.Profile, len(*addresses))

		for i, address := range *addresses {
			wg.Add(1)

			go func(i int, address string) {
				defer wg.Done()

				profile, err := getProfile(ctx, &GetProfileInput{
					Address: address,
				})

				if err != nil {
					logger.Errf(ctx, err, "err hydrating top profiles %v")
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

		cacheGateway.SaveTopViewedProfiles(ctx, &trimmedProfiles)

		return &trimmedProfiles
	}
}
