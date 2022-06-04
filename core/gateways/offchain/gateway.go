package offchain

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/core/gateways"
	"github.com/etheralley/etheralley-core-api/core/settings"
)

type gateway struct {
	settings           settings.ISettings
	logger             common.ILogger
	httpClient         common.IHttpClient
	cryptoPunkMetadata *cryptoPunksMetadata
	tokenMetadata      *map[string]tokenMetadata
}

func NewGateway(logger common.ILogger, settings settings.ISettings, httpClient common.IHttpClient) gateways.IOffchainGateway {
	return &gateway{
		settings,
		logger,
		httpClient,
		nil,
		&map[string]tokenMetadata{},
	}
}

func (gw *gateway) Init(ctx context.Context) error {
	if err := gw.initPunkMetadata(); err != nil {
		return err
	}
	if err := gw.initTokenMetadata(); err != nil {
		return err
	}
	return nil
}
