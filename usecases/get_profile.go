package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/mongo"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
)

func NewGetProfileUseCase(
	logger *common.Logger,
	cacheGateway *redis.Gateway,
	databaseGateway *mongo.Gateway,
	getDefaultProfile IGetDefaultProfileUseCase,
	getAllNonFungibleTokens IGetAllNonFungibleTokensUseCase,
	getAllFungibleTokens IGetAllFungibleTokensUseCase,
	getAllStatistics IGetAllStatisticsUseCase,
) IGetProfileUseCase {
	return GetProfile(logger, cacheGateway, databaseGateway, getDefaultProfile, getAllNonFungibleTokens, getAllFungibleTokens, getAllStatistics)
}

// first try to get the profile from the cache.
// if cache miss, go to database
// if database miss, fetch default tokens
// if database hit, re-fetch transient token info
func GetProfile(
	logger *common.Logger,
	cacheGateway gateways.ICacheGateway,
	databaseGateway gateways.IDatabaseGateway,
	getDefaultProfile IGetDefaultProfileUseCase,
	getAllNonFungibleTokens IGetAllNonFungibleTokensUseCase,
	getAllFungibleTokens IGetAllFungibleTokensUseCase,
	getAllStatistics IGetAllStatisticsUseCase,
) IGetProfileUseCase {
	return func(ctx context.Context, address string) (*entities.Profile, error) {
		if err := common.ValidateField(address, `required,eth_addr`); err != nil {
			return nil, err
		}

		profile, err := cacheGateway.GetProfileByAddress(ctx, address)

		if err == nil {
			logger.Debugf("cache hit for profile %v", address)
			return profile, nil
		}

		logger.Debugf("cache miss for profile %v", address)

		profile, err = databaseGateway.GetProfileByAddress(ctx, address)

		if err == common.ErrNotFound {
			logger.Debugf("db miss for profile %v", address)

			profile, err := getDefaultProfile(ctx, address)

			if err != nil {
				logger.Err(err, "err getting default profile")
				return nil, err
			}

			cacheGateway.SaveProfile(ctx, profile)

			return profile, nil
		}

		if err != nil {
			logger.Err(err, "err getting profile from db")
			return nil, err
		}

		logger.Debugf("db hit for profile %v", address)

		var wg sync.WaitGroup
		wg.Add(3)

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
			contracts := []entities.Contract{}
			for _, stats := range *profile.Statistics {
				contracts = append(contracts, *stats.Contract)
			}
			profile.Statistics = getAllStatistics(ctx, profile.Address, &contracts)
		}()

		wg.Wait()

		cacheGateway.SaveProfile(ctx, profile)

		return profile, nil
	}
}
