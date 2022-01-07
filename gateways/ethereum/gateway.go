package ethereum

import (
	"errors"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Gateway struct {
	settings *cmn.Settings
	logger   *cmn.Logger
}

func NewGateway(logger *cmn.Logger, settings *cmn.Settings) *Gateway {
	return &Gateway{
		settings,
		logger,
	}
}

func (gw *Gateway) getClient(blockchain string) (*ethclient.Client, error) {
	switch blockchain {
	case cmn.ETHEREUM:
		return ethclient.Dial(gw.settings.EthereumURI)
	case cmn.POLYGON:
		return ethclient.Dial(gw.settings.PolygonURI)
	case cmn.OPTIMISM:
		return ethclient.Dial(gw.settings.OptimismURI)
	case cmn.ARBITRUM:
		return ethclient.Dial(gw.settings.ArbitrumURI)
	}
	return nil, errors.New("invalid blockchain provided")
}

var zeroAddress = common.HexToAddress("0")
