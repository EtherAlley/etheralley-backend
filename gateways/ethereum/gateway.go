package ethereum

import (
	"fmt"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Gateway struct {
	settings *common.Settings
	logger   *common.Logger
	client   *ethclient.Client
}

func NewGateway(logger *common.Logger, settings *common.Settings) *Gateway {
	client, err := ethclient.Dial(fmt.Sprintf("https://%v.infura.io/v3/%v", settings.InfuraChain, settings.InfuraProjectId))

	if err != nil {
		panic(err)
	}

	return &Gateway{
		settings,
		logger,
		client,
	}
}
