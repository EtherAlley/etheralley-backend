package thegraph

import (
	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type gateway struct {
	settings    common.ISettings
	logger      common.ILogger
	graphClient common.IGraphQLClient
	httpClient  common.IHttpClient
}

func NewGateway(logger common.ILogger, settings common.ISettings, graphClient common.IGraphQLClient, httpClient common.IHttpClient) gateways.IBlockchainIndexGateway {
	return &gateway{
		settings,
		logger,
		graphClient,
		httpClient,
	}
}
