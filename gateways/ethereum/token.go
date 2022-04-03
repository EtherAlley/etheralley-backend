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

	balance, err := cmn.FunctionRetrier[*big.Int](ctx, gw.logger, instance.BalanceOf, &bind.CallOpts{}, adr)

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

	name, err = cmn.FunctionRetrier[string](ctx, gw.logger, instance.Name, &bind.CallOpts{})

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

	symbol, err = cmn.FunctionRetrier[string](ctx, gw.logger, instance.Symbol, &bind.CallOpts{})

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

	decimals, err = cmn.FunctionRetrier[uint8](ctx, gw.logger, instance.Decimals, &bind.CallOpts{})

	return
}
