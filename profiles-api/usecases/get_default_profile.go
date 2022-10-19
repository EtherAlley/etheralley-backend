package usecases

import (
	"context"
	"math/big"
	"sync"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
	"github.com/etheralley/etheralley-backend/profiles-api/gateways"
	"github.com/etheralley/etheralley-backend/profiles-api/settings"
)

func NewGetDefaultProfile(
	logger common.ILogger,
	settings settings.ISettings,
	blockchainGateway gateways.IBlockchainGateway,
	blochchainIndexGateway gateways.IBlockchainIndexGateway,
	offchainGateway gateways.IOffchainGateway,
	getAllFungibleTokens IGetAllFungibleTokensUseCase,
	getAllStatistics IGetAllStatisticsUseCase,
	getAllCurrencies IGetAllCurrenciesUseCase,
	resolveENSName IResolveENSNameUseCase,
) IGetDefaultProfileUseCase {
	return &getDefaultProfileUseCase{
		logger,
		settings,
		blockchainGateway,
		blochchainIndexGateway,
		offchainGateway,
		getAllFungibleTokens,
		getAllStatistics,
		getAllCurrencies,
		resolveENSName,
	}
}

type getDefaultProfileUseCase struct {
	logger                 common.ILogger
	settings               settings.ISettings
	blockchainGateway      gateways.IBlockchainGateway
	blochchainIndexGateway gateways.IBlockchainIndexGateway
	offchainGateway        gateways.IOffchainGateway
	getAllFungibleTokens   IGetAllFungibleTokensUseCase
	getAllStatistics       IGetAllStatisticsUseCase
	getAllCurrencies       IGetAllCurrenciesUseCase
	resolveENSName         IResolveENSNameUseCase
}

type IGetDefaultProfileUseCase interface {
	// Get a default profile for the provided address
	Do(ctx context.Context, input *GetDefaultProfileInput) (*entities.Profile, error)
}

type GetDefaultProfileInput struct {
	Address string `validate:"required,eth_addr"`
}

// Attempt to provide a pleasant default profile when none has been configured.
// Fetch nfts and tokens from alchemy and stats from the graph.
// Fetch primary ens name for address if configured
func (uc *getDefaultProfileUseCase) Do(ctx context.Context, input *GetDefaultProfileInput) (*entities.Profile, error) {
	if err := common.ValidateStruct(input); err != nil {
		return nil, err
	}

	profile := &entities.Profile{
		Address: input.Address,
		StoreAssets: &entities.StoreAssets{
			Premium:    false,
			BetaTester: false,
		},
		Interactions: &[]entities.Interaction{},
		Banned:       false,
	}
	var wg sync.WaitGroup
	wg.Add(6)

	go func() {
		defer wg.Done()

		name, err := uc.resolveENSName.Do(ctx, &ResolveENSNameInput{
			Address: input.Address,
		})

		if err == nil {
			profile.ENSName = name
		}
	}()

	go func() {
		defer wg.Done()

		nfts, err := uc.offchainGateway.GetNonFungibleTokens(ctx, input.Address)

		if err != nil {
			uc.logger.Warn(ctx).Err(err).Msgf("err fetching default nfts for address %v", input.Address)
			profile.NonFungibleTokens = &[]entities.NonFungibleToken{}
			return
		}

		trimmedNFTs := *nfts
		cutoff := len(trimmedNFTs)
		if cutoff > int(common.DEFAULT_NFT_CUTOFF) {
			cutoff = int(common.DEFAULT_NFT_CUTOFF)
		}
		trimmedNFTs = trimmedNFTs[:cutoff]

		profile.NonFungibleTokens = &trimmedNFTs
	}()

	go func() {
		defer wg.Done()

		contracts, err := uc.offchainGateway.GetFungibleContracts(ctx, input.Address)

		tokens := []GetFungibleTokenInput{}
		if err != nil || len(*contracts) == 0 {
			uc.logger.Debug(ctx).Err(err).Msgf("using default tokens for address %v. found %v contracts", input.Address, len(*contracts))
			for _, address := range uc.settings.DefaultTokenAddresses() {
				tokens = append(tokens, GetFungibleTokenInput{
					Address: input.Address,
					Token: &FungibleTokenInput{
						Contract: &ContractInput{
							Blockchain: common.ETHEREUM,
							Address:    address,
							Interface:  common.ERC20,
						},
					},
				})
			}
		} else {
			for _, contract := range *contracts {
				tokens = append(tokens, GetFungibleTokenInput{
					Address: input.Address,
					Token: &FungibleTokenInput{
						Contract: &ContractInput{
							Blockchain: contract.Blockchain,
							Address:    contract.Address,
							Interface:  contract.Interface,
						},
					},
				})
			}
		}

		cutoff := len(tokens)
		if cutoff > int(common.DEFAULT_TOKEN_CUTOFF) {
			cutoff = int(common.DEFAULT_TOKEN_CUTOFF)
		}
		trimmedTokens := tokens[:cutoff]

		profile.FungibleTokens = uc.getAllFungibleTokens.Do(ctx, &GetAllFungibleTokensInput{
			Tokens: &trimmedTokens,
		})
	}()

	go func() {
		defer wg.Done()

		stats := []GetStatisticsInput{
			{
				Address: input.Address,
				Statistic: &StatisticInput{
					Type: common.SWAP,
					Contract: &ContractInput{
						Address:    common.ZERO_ADDRESS,
						Interface:  common.UNISWAP_V2_EXCHANGE,
						Blockchain: common.ETHEREUM,
					},
				},
			},
			{
				Address: input.Address,
				Statistic: &StatisticInput{
					Type: common.SWAP,
					Contract: &ContractInput{
						Address:    common.ZERO_ADDRESS,
						Interface:  common.SUSHISWAP_EXCHANGE,
						Blockchain: common.ETHEREUM,
					},
				},
			},
			{
				Address: input.Address,
				Statistic: &StatisticInput{
					Type: common.STAKE,
					Contract: &ContractInput{
						Address:    common.ZERO_ADDRESS,
						Interface:  common.ROCKET_POOL,
						Blockchain: common.ETHEREUM,
					},
				},
			},
		}
		profile.Statistics = uc.getAllStatistics.Do(ctx, &GetAllStatisticsInput{
			Stats: &stats,
		})
	}()

	go func() {
		defer wg.Done()

		currencies := []GetCurrencyInput{}
		for _, chain := range []string{common.ETHEREUM, common.POLYGON, common.ARBITRUM, common.OPTIMISM} {
			currencies = append(currencies, GetCurrencyInput{
				Address:    input.Address,
				Blockchain: chain,
			})
		}
		profile.Currencies = uc.getAllCurrencies.Do(ctx, &GetAllCurrenciesInput{
			Currencies: &currencies,
		})
	}()

	go func() {
		defer wg.Done()

		balances, err := uc.blockchainGateway.GetStoreBalanceBatch(ctx, input.Address, &[]string{common.STORE_PREMIUM, common.STORE_BETA_TESTER})

		if err != nil {
			uc.logger.Info(ctx).Err(err).Msgf("err fetching store asset balance for %v", input.Address)
			profile.StoreAssets = &entities.StoreAssets{
				Premium:    false,
				BetaTester: false,
			}
			return
		}

		profile.StoreAssets = &entities.StoreAssets{
			Premium:    balances[0].Cmp(big.NewInt(0)) == 1,
			BetaTester: balances[1].Cmp(big.NewInt(0)) == 1,
		}
	}()

	wg.Wait()

	return profile, nil
}
