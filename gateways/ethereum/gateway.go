package ethereum

import "github.com/ethereum/go-ethereum/ethclient"

type Gateway struct {
	client *ethclient.Client
}

func NewGateway() *Gateway {
	client, err := ethclient.Dial("https://mainnet.infura.io")

	if err != nil {
		panic(err)
	}

	return &Gateway{
		client,
	}
}
