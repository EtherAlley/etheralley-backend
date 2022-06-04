package usecases

import (
	"context"
	"errors"
	"math/big"
	"sync"

	"github.com/etheralley/etheralley-apis/common"
	"github.com/etheralley/etheralley-apis/core/entities"
	"github.com/etheralley/etheralley-apis/core/gateways"
)

func NewGetProfile(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	cacheGateway gateways.ICacheGateway,
	databaseGateway gateways.IDatabaseGateway,
	getDefaultProfile IGetDefaultProfileUseCase,
	getAllNonFungibleTokens IGetAllNonFungibleTokensUseCase,
	getAllFungibleTokens IGetAllFungibleTokensUseCase,
	getAllStatistics IGetAllStatisticsUseCase,
	getAllCurrencies IGetAllCurrenciesUseCase,
	resolveENSName IResolveENSNameUseCase,
) IGetProfileUseCase {
	return &getProfileUseCase{
		logger,
		blockchainGateway,
		cacheGateway,
		databaseGateway,
		getDefaultProfile,
		getAllNonFungibleTokens,
		getAllFungibleTokens,
		getAllStatistics,
		getAllCurrencies,
		resolveENSName,
	}
}

type getProfileUseCase struct {
	logger                  common.ILogger
	blockchainGateway       gateways.IBlockchainGateway
	cacheGateway            gateways.ICacheGateway
	databaseGateway         gateways.IDatabaseGateway
	getDefaultProfile       IGetDefaultProfileUseCase
	getAllNonFungibleTokens IGetAllNonFungibleTokensUseCase
	getAllFungibleTokens    IGetAllFungibleTokensUseCase
	getAllStatistics        IGetAllStatisticsUseCase
	getAllCurrencies        IGetAllCurrenciesUseCase
	resolveENSName          IResolveENSNameUseCase
}

type IGetProfileUseCase interface {
	// get the profile for the provided address
	Do(ctx context.Context, input *GetProfileInput) (*entities.Profile, error)
}

type GetProfileInput struct {
	Address string `validate:"required,eth_addr"`
}

// First try to get the profile from the cache.
// If cache miss, go to database.
// If database miss, fetch default profile.
// If database hit, re-fetch transient info.
func (uc *getProfileUseCase) Do(ctx context.Context, input *GetProfileInput) (*entities.Profile, error) {
	if err := common.ValidateStruct(input); err != nil {
		return nil, err
	}

	uc.logger.Debug(ctx).Msgf("getting profile %v", input.Address)

	profile, err := uc.cacheGateway.GetProfileByAddress(ctx, input.Address)

	if err == nil {
		uc.logger.Debug(ctx).Msgf("cache hit %v", input.Address)
		return profile, nil
	}

	uc.logger.Debug(ctx).Msgf("cache miss %v", input.Address)

	profile, err = uc.databaseGateway.GetProfileByAddress(ctx, input.Address)

	if errors.Is(err, common.ErrNotFound) {
		uc.logger.Debug(ctx).Msgf("db miss %v", input.Address)

		profile, err := uc.getDefaultProfile.Do(ctx, &GetDefaultProfileInput{
			Address: input.Address,
		})

		if err != nil {
			uc.logger.Warn(ctx).Err(err).Msgf("err getting default profile %v", input.Address)
			return nil, err
		}

		uc.cacheGateway.SaveProfile(ctx, profile)

		return profile, nil
	}

	if err != nil {
		uc.logger.Warn(ctx).Err(err).Msgf("err getting db profile %v", input.Address)
		return nil, err
	}

	uc.logger.Debug(ctx).Msgf("db hit %v", input.Address)

	var wg sync.WaitGroup
	wg.Add(6)

	go func() {
		defer wg.Done()

		// Not all addresses have an ens name. We should not propigate an error for this
		name, _ := uc.resolveENSName.Do(ctx, &ResolveENSNameInput{
			Address: profile.Address,
		})

		profile.ENSName = name
	}()

	go func() {
		defer wg.Done()

		nfts := []GetNonFungibleTokenInput{}
		for _, nft := range *profile.NonFungibleTokens {
			nfts = append(nfts, GetNonFungibleTokenInput{
				Address: profile.Address,
				NonFungibleToken: &NonFungibleTokenInput{
					TokenId: nft.TokenId,
					Contract: &ContractInput{
						Blockchain: nft.Contract.Blockchain,
						Address:    nft.Contract.Address,
						Interface:  nft.Contract.Interface,
					},
				},
			})
		}

		profile.NonFungibleTokens = uc.getAllNonFungibleTokens.Do(ctx, &GetAllNonFungibleTokensInput{
			NonFungibleTokens: &nfts,
		})
	}()

	go func() {
		defer wg.Done()

		tokens := []GetFungibleTokenInput{}
		for _, token := range *profile.FungibleTokens {
			tokens = append(tokens, GetFungibleTokenInput{
				Address: profile.Address,
				Token: &FungibleTokenInput{
					Contract: &ContractInput{
						Blockchain: token.Contract.Blockchain,
						Address:    token.Contract.Address,
						Interface:  token.Contract.Interface,
					},
				},
			})
		}

		profile.FungibleTokens = uc.getAllFungibleTokens.Do(ctx, &GetAllFungibleTokensInput{
			Tokens: &tokens,
		})
	}()

	go func() {
		defer wg.Done()

		stats := []GetStatisticsInput{}
		for _, stat := range *profile.Statistics {
			stats = append(stats, GetStatisticsInput{
				Address: profile.Address,
				Statistic: &StatisticInput{
					Type: stat.Type,
					Contract: &ContractInput{
						Blockchain: stat.Contract.Blockchain,
						Address:    stat.Contract.Address,
						Interface:  stat.Contract.Interface,
					},
				},
			})
		}

		profile.Statistics = uc.getAllStatistics.Do(ctx, &GetAllStatisticsInput{
			Stats: &stats,
		})
	}()

	go func() {
		defer wg.Done()

		currencies := []GetCurrencyInput{}
		for _, currency := range *profile.Currencies {
			currencies = append(currencies, GetCurrencyInput{
				Address:    profile.Address,
				Blockchain: currency.Blockchain,
			})
		}

		profile.Currencies = uc.getAllCurrencies.Do(ctx, &GetAllCurrenciesInput{
			Currencies: &currencies,
		})
	}()

	go func() {
		defer wg.Done()

		profile.StoreAssets = &entities.StoreAssets{}

		if balances, err := uc.blockchainGateway.GetStoreBalanceBatch(ctx, input.Address, &[]string{common.STORE_PREMIUM, common.STORE_BETA_TESTER}); err == nil {
			profile.StoreAssets.Premium = balances[0].Cmp(big.NewInt(0)) == 1
			profile.StoreAssets.BetaTester = balances[1].Cmp(big.NewInt(0)) == 1
		}
	}()

	wg.Wait()

	uc.cacheGateway.SaveProfile(ctx, profile)

	return profile, nil
}
