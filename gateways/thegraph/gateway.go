package thegraph

import (
	"errors"
	"fmt"

	"github.com/etheralley/etheralley-core-api/common"
)

type Gateway struct {
	settings *common.Settings
	logger   *common.Logger
	client   *common.GraphQLClient
}

func NewGateway(logger *common.Logger, settings *common.Settings, client *common.GraphQLClient) *Gateway {
	return &Gateway{
		settings,
		logger,
		client,
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
		}
	case common.POLYGON:
		switch i {
		case common.SUSHISWAP_EXCHANGE:
			return fmt.Sprintf("%v/sushiswap/matic-exchange", gw.settings.TheGraphHostedURI), nil
		}
	}

	return "", errors.New("unsupported interface blockchain combination")
}
