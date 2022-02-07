package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

// first try to get the profile from the cache.
// if cache miss, go to database
// if database miss, fetch default profile
// if database hit, re-fetch transient token info and ens name
func NewGetProfile(
	logger common.ILogger,
	cacheGateway gateways.ICacheGateway,
	databaseGateway gateways.IDatabaseGateway,
	getDefaultProfile IGetDefaultProfileUseCase,
	getAllNonFungibleTokens IGetAllNonFungibleTokensUseCase,
	getAllFungibleTokens IGetAllFungibleTokensUseCase,
	getAllStatistics IGetAllStatisticsUseCase,
	resolveENSName IResolveENSNameUseCase,
) IGetProfileUseCase {
	return func(ctx context.Context, address string) (*entities.Profile, error) {
		if err := common.ValidateField(address, `required,eth_addr`); err != nil {
			return nil, err
		}

		profile, err := cacheGateway.GetProfileByAddress(ctx, address)

		if err == nil {
			logger.Debugf(ctx, "cache hit for profile %v", address)
			return profile, nil
		}

		logger.Debugf(ctx, "cache miss for profile %v", address)

		profile, err = databaseGateway.GetProfileByAddress(ctx, address)

		if err == common.ErrNotFound {
			logger.Debugf(ctx, "db miss for profile %v", address)

			profile, err := getDefaultProfile(ctx, address)

			if err != nil {
				logger.Err(ctx, err, "err getting default profile")
				return nil, err
			}

			cacheGateway.SaveProfile(ctx, profile)

			return profile, nil
		}

		if err != nil {
			logger.Err(ctx, err, "err getting profile from db")
			return nil, err
		}

		logger.Debugf(ctx, "db hit for profile %v", address)

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

		return profile, nil
	}
}
