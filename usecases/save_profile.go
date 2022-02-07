package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

// fetch transient info for nfts/tokens/stats/etc being submitted
// try to save the profile to the cache
// regardless of error, save the profile to the database
func NewSaveProfile(
	logger common.ILogger,
	cacheGateway gateways.ICacheGateway,
	databaseGateway gateways.IDatabaseGateway,
	getAllNonFungibleTokens IGetAllNonFungibleTokensUseCase,
	getAllFungibleTokens IGetAllFungibleTokensUseCase,
	getAllStatistics IGetAllStatisticsUseCase,
	resolveENSName IResolveENSNameUseCase,
) ISaveProfileUseCase {
	return func(ctx context.Context, address string, profile *entities.Profile) error {
		if err := common.ValidateStruct(profile); err != nil {
			return err
		}

		profile.Address = address

		var wg sync.WaitGroup
		wg.Add(4)

		go func() {
			defer wg.Done()
			profile.NonFungibleTokens = getAllNonFungibleTokens(ctx, profile.Address, profile.NonFungibleTokens)
		}()

		go func() {
			defer wg.Done()
			contracts := []entities.Contract{}
			for _, token := range *profile.FungibleTokens {
				contracts = append(contracts, *token.Contract)
			}
			profile.FungibleTokens = getAllFungibleTokens(ctx, profile.Address, &contracts)
		}()

		go func() {
			defer wg.Done()
			input := GetAllStatisticsInput{
				Address: profile.Address,
				Stats:   &[]StatisticInput{},
			}
			for _, stats := range *profile.Statistics {
				*input.Stats = append(*input.Stats, StatisticInput{
					Contract: stats.Contract,
					Type:     stats.Type,
				})
			}
			profile.Statistics = getAllStatistics(ctx, &input)
		}()

		go func() {
			defer wg.Done()
			name, err := resolveENSName(ctx, address)
			if err != nil {
				profile.ENSName = "" // Not all addresses have an ens name. We should not propigate an erro for this
			} else {
				profile.ENSName = name
			}
		}()

		wg.Wait()

		cacheGateway.SaveProfile(ctx, profile)

		return databaseGateway.SaveProfile(ctx, profile)
	}
}
