package thegraph

import (
	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type gateway struct {
	settings    common.ISettings
	logger      common.ILogger
	graphClient common.IGraphQLClient
}

func NewGateway(logger common.ILogger, settings common.ISettings, graphClient common.IGraphQLClient) gateways.IBlockchainIndexGateway {
	return &gateway{
		settings,
		logger,
		graphClient,
	}
}
