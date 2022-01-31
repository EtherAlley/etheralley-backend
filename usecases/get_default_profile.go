package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/thegraph"
)

func NewGetDefaultProfileUseCase(
	logger *common.Logger,
	settings *common.Settings,
	blochchainIndexGateway *thegraph.Gateway,
	getAllFungibleTokens IGetAllFungibleTokensUseCase,
	getAllStatistics IGetAllStatisticsUseCase,
) IGetDefaultProfileUseCase {
	return GetDefaultProfile(logger, settings, blochchainIndexGateway, getAllFungibleTokens, getAllStatistics)
}

// attempt to provide a pleasent default profile when none has been configured.
// fetch nfts and stats from the graph and fetch tokens from a fixed list.
func GetDefaultProfile(
	logger *common.Logger,
	settings *common.Settings,
	blochchainIndexGateway gateways.IBlockchainIndexGateway,
	getAllFungibleTokens IGetAllFungibleTokensUseCase,
	getAllStatistics IGetAllStatisticsUseCase,
) IGetDefaultProfileUseCase {
	return func(ctx context.Context, address string) (*entities.Profile, error) {
		var nfts *[]entities.NonFungibleToken
		var tokens *[]entities.FungibleToken
		var stats *[]entities.Statistic
		var wg sync.WaitGroup
		wg.Add(3)

		go func() {
			defer wg.Done()
			nfts = blochchainIndexGateway.GetNonFungibleTokens(ctx, address)
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

			tokens = getAllFungibleTokens(ctx, address, &contracts)
		}()

		go func() {
			defer wg.Done()
			contracts := []entities.Contract{
				{
					Address:    common.ZERO_ADDRESS,
					Interface:  common.UNISWAP_V2_EXCHANGE,
					Blockchain: common.ETHEREUM,
				},
				{
					Address:    common.ZERO_ADDRESS,
					Interface:  common.SUSHISWAP_EXCHANGE,
					Blockchain: common.ETHEREUM,
				},
			}
			stats = getAllStatistics(ctx, address, &contracts)
		}()

		wg.Wait()

		profile := &entities.Profile{
			Address:           address,
			NonFungibleTokens: nfts,
			FungibleTokens:    tokens,
			Statistics:        stats,
		}

		return profile, nil
	}
}
