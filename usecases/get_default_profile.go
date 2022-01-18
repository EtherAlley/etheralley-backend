package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/opensea"
)

func NewGetDefaultProfileUseCase(
	logger *common.Logger,
	settings *common.Settings,
	nftApiGateway *opensea.Gateway,
	getAllFungibleTokens GetAllFungibleTokensUseCase,
) GetDefaultProfileUseCase {
	return GetDefaultProfile(logger, settings, nftApiGateway, getAllFungibleTokens)
}

// attempt to provide a pleasent default profile when none has been configured.
// fetch nfts from opensea and fetch tokens from a fixed list.
func GetDefaultProfile(
	logger *common.Logger,
	settings *common.Settings,
	nftApiGateway gateways.INonFungibleAPIGateway,
	getAllFungibleTokens GetAllFungibleTokensUseCase,
) GetDefaultProfileUseCase {
	return func(ctx context.Context, address string) (*entities.Profile, error) {
		var nfts *[]entities.NonFungibleToken
		var tokens *[]entities.FungibleToken

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			nfts = nftApiGateway.GetNonFungibleTokens(address)
		}()

		go func() {
			defer wg.Done()

			var knownContracts []string
			if settings.IsDev() {
				knownContracts = KnownGoerliContracts
			} else {
				knownContracts = KnownMainnetContracts
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

		wg.Wait()

		profile := &entities.Profile{
			Address:           address,
			NonFungibleTokens: nfts,
			FungibleTokens:    tokens,
		}

		return profile, nil
	}
}

var KnownGoerliContracts = []string{
	common.UNI_GOERLI,
	common.LINK_GOERLI,
	common.HEX_GOERLI,
	common.DAI_GOERLI,
	common.BUSD_GOERLI,
	common.USDC_GOERLI,
	common.USDT_GOERLI,
	common.WETH_GOERLI,
}

var KnownMainnetContracts = []string{
	common.USDT_MAINNET,
	common.BNB_MAINNET,
	common.USDC_MAINNET,
	common.HEX_MAINNET,
	common.MATIC_MAINNET,
	common.SHIB_MAINNET,
	common.BUSD_MAINNET,
	common.LINK_MAINNET,
	common.CRO_MAINNET,
	common.WBTC_MAINNET,
	common.UST_MAINNET,
	common.DAI_MAINNET,
	common.UNI_MAINNET,
}
