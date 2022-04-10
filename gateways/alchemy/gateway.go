package alchemy

import (
	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type gateway struct {
	settings   common.ISettings
	logger     common.ILogger
	httpClient common.IHttpClient
}

func NewGateway(logger common.ILogger, settings common.ISettings, httpClient common.IHttpClient) gateways.INFTAPIGateway {
	return &gateway{
		settings,
		logger,
		httpClient,
	}
}
