package thegraph

import (
	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/gateways"
	"github.com/etheralley/etheralley-backend/profiles-api/settings"
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
