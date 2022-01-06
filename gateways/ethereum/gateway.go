package ethereum

import (
	"fmt"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Gateway struct {
	settings *cmn.Settings
	logger   *cmn.Logger
	client   *ethclient.Client
}

func NewGateway(logger *cmn.Logger, settings *cmn.Settings) *Gateway {
	var blockchain string
	switch settings.EthereumNetwork {
	case "testnet":
		blockchain = "goerli"
	case "mainnet":
		blockchain = "mainnet"
	}
	client, err := ethclient.Dial(fmt.Sprintf("https://eth-%v.alchemyapi.io/v2/%v", blockchain, settings.EthereumAPIKey))

	if err != nil {
		panic(err)
	}

	return &Gateway{
		settings,
		logger,
		client,
	}
}

var zeroAddress = common.HexToAddress("0")
