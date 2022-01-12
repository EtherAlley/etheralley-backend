package ethereum

import (
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (gw *Gateway) GetFungibleBalance(address string, contract *entities.Contract) (string, error) {
	client, err := gw.getClient(contract.Blockchain)

	if err != nil {
		return "", err
	}

	contractAddress := common.HexToAddress(contract.Address)
	adr := common.HexToAddress(address)

	instance, err := contracts.NewErc20(contractAddress, client)

	if err != nil {
		return "", err
	}

	balance, err := instance.BalanceOf(&bind.CallOpts{}, adr)

	if err != nil {
		return "", err
	}

	return balance.String(), err
}

func (gw *Gateway) GetFungibleName(contract *entities.Contract) (name string, err error) {
	client, err := gw.getClient(contract.Blockchain)

	if err != nil {
		return
	}

	contractAddress := common.HexToAddress(contract.Address)

	instance, err := contracts.NewErc20(contractAddress, client)

	if err != nil {
		return
	}

	name, err = instance.Name(&bind.CallOpts{})

	return
}

func (gw *Gateway) GetFungibleSymbol(contract *entities.Contract) (symbol string, err error) {
	client, err := gw.getClient(contract.Blockchain)

	if err != nil {
		return
	}

	contractAddress := common.HexToAddress(contract.Address)

	instance, err := contracts.NewErc20(contractAddress, client)

	if err != nil {
		return
	}

	symbol, err = instance.Symbol(&bind.CallOpts{})

	return
}

func (gw *Gateway) GetFungibleDecimals(contract *entities.Contract) (decimals uint8, err error) {
	client, err := gw.getClient(contract.Blockchain)

	if err != nil {
		return
	}

	contractAddress := common.HexToAddress(contract.Address)

	instance, err := contracts.NewErc20(contractAddress, client)

	if err != nil {
		return
	}

	decimals, err = instance.Decimals(&bind.CallOpts{})

	return
}
