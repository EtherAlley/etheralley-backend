package usecases

import (
	"context"
	"math/big"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type GetDefaultProfileInput struct {
	Address string `validate:"required,eth_addr"`
}

// Get a default profile for the provided address
type IGetDefaultProfileUseCase func(ctx context.Context, input *GetDefaultProfileInput) (*entities.Profile, error)

// Attempt to provide a pleasent default profile when none has been configured.
//
// Fetch nfts and stats from the graph and fetch tokens from a fixed list.
//
// Fetch primary ens name for address if configured
func NewGetDefaultProfile(
	logger common.ILogger,
	settings common.ISettings,
	blockchainGateway gateways.IBlockchainGateway,
	blochchainIndexGateway gateways.IBlockchainIndexGateway,
	nftApiGateway gateways.INFTAPIGateway,
	getAllFungibleTokens IGetAllFungibleTokensUseCase,
	getAllStatistics IGetAllStatisticsUseCase,
	resolveENSName IResolveENSNameUseCase,
) IGetDefaultProfileUseCase {
	return func(ctx context.Context, input *GetDefaultProfileInput) (*entities.Profile, error) {
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
		}
		var wg sync.WaitGroup
		wg.Add(5)

		go func() {
			defer wg.Done()

			// Not all addresses have an ens name. We should not propigate an error for this
			name, _ := resolveENSName(ctx, &ResolveENSNameInput{
				Address: input.Address,
			})

			profile.ENSName = name
		}()

		go func() {
			defer wg.Done()
			profile.NonFungibleTokens = nftApiGateway.GetNonFungibleTokens(ctx, input.Address)
		}()

		go func() {
			defer wg.Done()

			tokens := []GetFungibleTokenInput{}
			for _, address := range settings.DefaultTokenAddresses() {
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

			profile.FungibleTokens = getAllFungibleTokens(ctx, &GetAllFungibleTokensInput{
				Tokens: &tokens,
			})
		}()

		go func() {
			defer wg.Done()

			stats := []GetStatisticsInput{}
			for _, intf := range []string{common.UNISWAP_V2_EXCHANGE, common.SUSHISWAP_EXCHANGE} {
				stats = append(stats, GetStatisticsInput{
					Address: input.Address,
					Statistic: &StatisticInput{
						Type: common.SWAP,
						Contract: &ContractInput{
							Address:    common.ZERO_ADDRESS,
							Interface:  intf,
							Blockchain: common.ETHEREUM,
						},
					},
				})
			}
			profile.Statistics = getAllStatistics(ctx, &GetAllStatisticsInput{
				Stats: &stats,
			})
		}()

		go func() {
			defer wg.Done()

			profile.StoreAssets = &entities.StoreAssets{}

			if balances, err := blockchainGateway.GetStoreBalanceBatch(ctx, input.Address, &[]string{common.STORE_PREMIUM, common.STORE_BETA_TESTER}); err == nil {
				profile.StoreAssets.Premium = balances[0].Cmp(big.NewInt(0)) == 1
				profile.StoreAssets.BetaTester = balances[1].Cmp(big.NewInt(0)) == 1
			}
		}()

		wg.Wait()

		return profile, nil
	}
}
