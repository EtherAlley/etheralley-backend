package thegraph

import (
	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/core/gateways"
	"github.com/etheralley/etheralley-core-api/core/settings"
)

type gateway struct {
	settings    settings.ISettings
	logger      common.ILogger
	graphClient common.IGraphQLClient
}

func NewGateway(logger common.ILogger, settings settings.ISettings, graphClient common.IGraphQLClient) gateways.IBlockchainIndexGateway {
	return &gateway{
		settings,
		logger,
		graphClient,
	}
}
