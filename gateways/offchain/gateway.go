package offchain

import (
	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type gateway struct {
	settings           common.ISettings
	logger             common.ILogger
	httpClient         common.IHttpClient
	cryptoPunkMetadata *cryptoPunksMetadata
	tokenMetadata      *map[string]tokenMetadata
}

func NewGateway(logger common.ILogger, settings common.ISettings, httpClient common.IHttpClient) gateways.IOffchainGateway {
	return &gateway{
		settings,
		logger,
		httpClient,
		nil,
		&map[string]tokenMetadata{},
	}
}

func (gw *gateway) Init() {
	gw.initPunkMetadata()
	gw.initTokenMetadata()
}
