package ethereum

import (
	"context"
	"errors"
	"fmt"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type gateway struct {
	settings cmn.ISettings
	logger   cmn.ILogger
	http     cmn.IHttpClient
}

func NewGateway(logger cmn.ILogger, settings cmn.ISettings, http cmn.IHttpClient) gateways.IBlockchainGateway {
	return &gateway{
		settings,
		logger,
		http,
	}
}

func (gw *gateway) getClient(ctx context.Context, blockchain cmn.Blockchain) (*ethclient.Client, error) {
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

// Parse for go-ethereum http error to determine if its retryable.
// Wrap in common.ErrRetryable if status code is 429
// and add context
func tryWrapRetryable(context string, err error) error {
	if err == nil {
		return nil
	}

	var e rpc.HTTPError
	if errors.As(err, &e) && e.StatusCode == 429 {
		return fmt.Errorf("%v %v %w", context, err, cmn.ErrRetryable)
	}

	return err
}
