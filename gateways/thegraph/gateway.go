package thegraph

import (
	"errors"
	"fmt"

	"github.com/etheralley/etheralley-core-api/common"
)

type Gateway struct {
	settings    *common.Settings
	logger      *common.Logger
	graphClient *common.GraphQLClient
	httpClient  *common.HttpClient
}

func NewGateway(logger *common.Logger, settings *common.Settings, graphClient *common.GraphQLClient, httpClient *common.HttpClient) *Gateway {
	return &Gateway{
		settings,
		logger,
		graphClient,
		httpClient,
	}
}

func (gw Gateway) GetSubgraphUrl(b common.Blockchain, i common.Interface) (string, error) {
	switch b {
	case common.ETHEREUM:
		switch i {
		case common.SUSHISWAP_EXCHANGE:
			return fmt.Sprintf("%v/sushiswap/exchange", gw.settings.TheGraphHostedURI), nil
		case common.UNISWAP_V2_EXCHANGE:
			return fmt.Sprintf("%v/uniswap/uniswap-v2", gw.settings.TheGraphHostedURI), nil
		case common.UNISWAP_V3_EXCHANGE:
			return fmt.Sprintf("%v/uniswap/uniswap-v3", gw.settings.TheGraphHostedURI), nil
		case common.ERC721:
			return fmt.Sprintf("%v/0x7859821024e633c5dc8a4fcf86fc52e7720ce525-0", gw.settings.TheGraphURI), nil
		}
	case common.POLYGON:
		switch i {
		case common.SUSHISWAP_EXCHANGE:
			return fmt.Sprintf("%v/sushiswap/matic-exchange", gw.settings.TheGraphHostedURI), nil
		}
	}

	return "", errors.New("unsupported interface blockchain combination")
}
