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

func NewSaveProfileUseCase(
	logger *common.Logger,
	cacheGateway *redis.Gateway,
	databaseGateway *mongo.Gateway,
	getAllNonFungibleTokens GetAllNonFungibleTokensUseCase,
	getAllFungibleTokens GetAllFungibleTokensUseCase,
	getAllStatistics GetAllStatisticsUseCase,
) SaveProfileUseCase {
	return SaveProfile(
		logger,
		cacheGateway,
		databaseGateway,
		getAllNonFungibleTokens,
		getAllFungibleTokens,
		getAllStatistics,
	)
}

// fetch metadata and ownership of nfts being submitted
// try to save the profile to the cache
// regardless of error, save the profile to the database
func SaveProfile(
	logger *common.Logger,
	cacheGateway gateways.ICacheGateway,
	databaseGateway gateways.IDatabaseGateway,
	getAllNonFungibleTokens GetAllNonFungibleTokensUseCase,
	getAllFungibleTokens GetAllFungibleTokensUseCase,
	getAllStatistics GetAllStatisticsUseCase,
) SaveProfileUseCase {
	return func(ctx context.Context, profile *entities.Profile) error {
		if err := common.ValidateStruct(profile); err != nil {
			return err
		}

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

		return databaseGateway.SaveProfile(ctx, profile)
	}
}
