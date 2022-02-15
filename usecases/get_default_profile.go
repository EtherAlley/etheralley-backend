package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

// attempt to provide a pleasent default profile when none has been configured.
// fetch nfts and stats from the graph and fetch tokens from a fixed list.
// fetch primary ens name for address if configured
func NewGetDefaultProfile(
	logger common.ILogger,
	settings common.ISettings,
	blochchainIndexGateway gateways.IBlockchainIndexGateway,
	getAllFungibleTokens IGetAllFungibleTokensUseCase,
	getAllStatistics IGetAllStatisticsUseCase,
	resolveENSName IResolveENSNameUseCase,
) IGetDefaultProfileUseCase {
	return func(ctx context.Context, address string) (*entities.Profile, error) {
		if err := common.ValidateField(address, `required,eth_addr`); err != nil {
			return nil, err
		}

		profile := &entities.Profile{
			Address:      address,
			Interactions: &[]entities.Interaction{},
		}
		var wg sync.WaitGroup
		wg.Add(4)

		go func() {
			defer wg.Done()
			profile.NonFungibleTokens = blochchainIndexGateway.GetNonFungibleTokens(ctx, address)
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

			contracts := []entities.Contract{}
			for _, address := range knownContracts {
				contracts = append(contracts, entities.Contract{
					Address:    address,
					Blockchain: common.ETHEREUM,
					Interface:  common.ERC20,
				})
			}

			profile.FungibleTokens = getAllFungibleTokens(ctx, address, &contracts)
		}()

		go func() {
			defer wg.Done()
			input := GetAllStatisticsInput{
				Address: address,
				Stats: &[]StatisticInput{
					{
						Type: common.SWAP,
						Contract: &entities.Contract{
							Address:    common.ZERO_ADDRESS,
							Interface:  common.UNISWAP_V2_EXCHANGE,
							Blockchain: common.ETHEREUM,
						},
					},
					{
						Type: common.SWAP,
						Contract: &entities.Contract{
							Address:    common.ZERO_ADDRESS,
							Interface:  common.SUSHISWAP_EXCHANGE,
							Blockchain: common.ETHEREUM,
						},
					},
				},
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

		return profile, nil
	}
}
