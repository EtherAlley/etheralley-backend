package ethereum

import (
	"context"
	"errors"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Gateway struct {
	settings cmn.ISettings
	logger   cmn.ILogger
	http     cmn.IHttpClient
}

func NewGateway(logger cmn.ILogger, settings cmn.ISettings, http cmn.IHttpClient) gateways.IBlockchainGateway {
	return &Gateway{
		settings,
		logger,
		http,
	}
}

func (gw *Gateway) getClient(ctx context.Context, blockchain cmn.Blockchain) (*ethclient.Client, error) {
	switch blockchain {
	case cmn.ETHEREUM:
		return ethclient.DialContext(ctx, gw.settings.EthereumURI())
	case cmn.POLYGON:
		return ethclient.DialContext(ctx, gw.settings.PolygonURI())
	case cmn.OPTIMISM:
		return ethclient.DialContext(ctx, gw.settings.OptimismURI())
	case cmn.ARBITRUM:
		return ethclient.DialContext(ctx, gw.settings.ArbitrumURI())
	}
	return nil, errors.New("invalid blockchain provided")
}
