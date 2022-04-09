package ethereum

import (
	"context"
	"math/big"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (gw *gateway) GetFungibleBalance(ctx context.Context, address string, contract *entities.Contract) (string, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return "", err
	}

	contractAddress := common.HexToAddress(contract.Address)
	adr := common.HexToAddress(address)

	instance, err := contracts.NewErc20(contractAddress, client)

	if err != nil {
		return "", err
	}

	balance, err := cmn.FunctionRetrier(ctx, func() (*big.Int, error) {
		balance, err := instance.BalanceOf(&bind.CallOpts{}, adr)
		return balance, tryWrapRetryable("get erc20 balance", err)
	})

	if err != nil {
		return "", err
	}

	return balance.String(), err
}

func (gw *gateway) GetFungibleName(ctx context.Context, contract *entities.Contract) (name string, err error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return
	}

	contractAddress := common.HexToAddress(contract.Address)

	instance, err := contracts.NewErc20(contractAddress, client)

	if err != nil {
		return
	}

	name, err = cmn.FunctionRetrier(ctx, func() (string, error) {
		name, err := instance.Name(&bind.CallOpts{})
		return name, tryWrapRetryable("get erc20 name", err)
	})

	return
}

func (gw *gateway) GetFungibleSymbol(ctx context.Context, contract *entities.Contract) (symbol string, err error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return
	}

	contractAddress := common.HexToAddress(contract.Address)

	instance, err := contracts.NewErc20(contractAddress, client)

	if err != nil {
		return
	}

	symbol, err = cmn.FunctionRetrier(ctx, func() (string, error) {
		symbol, err := instance.Symbol(&bind.CallOpts{})
		return symbol, tryWrapRetryable("get erc20 symbol", err)
	})

	return
}

func (gw *gateway) GetFungibleDecimals(ctx context.Context, contract *entities.Contract) (decimals uint8, err error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return
	}

	contractAddress := common.HexToAddress(contract.Address)

	instance, err := contracts.NewErc20(contractAddress, client)

	if err != nil {
		return
	}

	decimals, err = cmn.FunctionRetrier(ctx, func() (uint8, error) {
		decimals, err := instance.Decimals(&bind.CallOpts{})
		return decimals, tryWrapRetryable("get erc20 decimals", err)
	})

	return
}
