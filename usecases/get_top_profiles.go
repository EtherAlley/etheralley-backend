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
	resolveENSName IResolveENSNameUseCase,
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
		profiles := make([]entities.Profile, len(*addresses))

		for i, address := range *addresses {
			wg.Add(1)

			go func(i int, address string) {
				defer wg.Done()

				ensName, _ := resolveENSName(ctx, &ResolveENSNameInput{
					Address: address,
				})

				profiles[i] = entities.Profile{
					Address: address,
					ENSName: ensName,
				}
			}(i, address)
		}

		wg.Wait()

		cacheGateway.SaveTopViewedProfiles(ctx, &profiles)

		return &profiles
	}
}
