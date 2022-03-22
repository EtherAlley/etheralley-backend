package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type GetDefaultProfileInput struct {
	Address string `validate:"required,eth_addr"`
}

// get a default profile for the provided address
type IGetDefaultProfileUseCase func(ctx context.Context, input *GetDefaultProfileInput) (*entities.Profile, error)

// attempt to provide a pleasent default profile when none has been configured.
// fetch nfts and stats from the graph and fetch tokens from a fixed list.
// fetch primary ens name for address if configured
func NewGetDefaultProfile(
	logger common.ILogger,
	settings common.ISettings,
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
			Address:      input.Address,
			Interactions: &[]entities.Interaction{},
		}
		var wg sync.WaitGroup
		wg.Add(4)

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

			var knownContracts []string
			if settings.IsDev() {
				knownContracts = []string{
					common.UNI_GOERLI,
					common.LINK_GOERLI,
					common.HEX_GOERLI,
					common.SHIB_GOERLI,
					common.DAI_GOERLI,
					common.CRO_GOERLI,
				}
			} else {
				knownContracts = []string{
					common.UNI_MAINNET,
					common.LINK_MAINNET,
					common.HEX_MAINNET,
					common.SHIB_MAINNET,
					common.DAI_MAINNET,
					common.CRO_MAINNET,
				}
			}

			tokens := []GetFungibleTokenInput{}
			for _, address := range knownContracts {
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
			input := GetAllStatisticsInput{
				Stats: &[]GetStatisticsInput{
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
				},
			}
			profile.Statistics = getAllStatistics(ctx, &input)
		}()

		wg.Wait()

		return profile, nil
	}
}
