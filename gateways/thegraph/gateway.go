package thegraph

import (
	"errors"
	"fmt"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type Gateway struct {
	settings    common.ISettings
	logger      common.ILogger
	graphClient common.IGraphQLClient
	httpClient  common.IHttpClient
}

func NewGateway(logger common.ILogger, settings common.ISettings, graphClient common.IGraphQLClient, httpClient common.IHttpClient) gateways.IBlockchainIndexGateway {
	return &Gateway{
		settings,
		logger,
		graphClient,
		httpClient,
	}
}

func (gw Gateway) GetSubgraphUrl(b common.Blockchain, i common.Interface) (string, error) {
	hostedURI := gw.settings.TheGraphHostedURI()
	decURI := gw.settings.TheGraphURI()
	switch b {
	case common.ETHEREUM:
		switch i {
		case common.SUSHISWAP_EXCHANGE:
			return fmt.Sprintf("%v/sushiswap/exchange", hostedURI), nil
		case common.UNISWAP_V2_EXCHANGE:
			return fmt.Sprintf("%v/uniswap/uniswap-v2", hostedURI), nil
		case common.UNISWAP_V3_EXCHANGE:
			return fmt.Sprintf("%v/uniswap/uniswap-v3", hostedURI), nil
		case common.ERC721:
			return fmt.Sprintf("%v/0x7859821024e633c5dc8a4fcf86fc52e7720ce525-0", decURI), nil
		}
	case common.POLYGON:
		switch i {
		case common.SUSHISWAP_EXCHANGE:
			return fmt.Sprintf("%v/sushiswap/matic-exchange", hostedURI), nil
		}
	case common.ARBITRUM:
		switch i {
		case common.SUSHISWAP_EXCHANGE:
			return fmt.Sprintf("%v/sushiswap/arbitrum-exchange", hostedURI), nil
		}
	}

	return "", errors.New("unsupported interface blockchain combination")
}
