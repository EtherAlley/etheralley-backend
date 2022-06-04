package ethereum

import (
	"context"
	"fmt"
	"math/big"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/core/entities"
	"github.com/etheralley/etheralley-core-api/core/gateways/ethereum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (gw *gateway) GetERC20Balance(ctx context.Context, address string, contract *entities.Contract) (string, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return "", fmt.Errorf("erc20 balance client %w", err)
	}

	contractAddress := common.HexToAddress(contract.Address)
	adr := common.HexToAddress(address)

	instance, err := contracts.NewErc20(contractAddress, client)

	if err != nil {
		return "", fmt.Errorf("erc20 balance contract %w", err)
	}

	balance, err := cmn.FunctionRetrier(ctx, func() (*big.Int, error) {
		balance, err := instance.BalanceOf(&bind.CallOpts{}, adr)
		return balance, gw.tryWrapRetryable(ctx, "erc20 balance retry", err)
	})

	if err != nil {
		return "", fmt.Errorf("erc20 balance %w", err)
	}

	return balance.String(), nil
}

func (gw *gateway) GetERC20Name(ctx context.Context, contract *entities.Contract) (string, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return "", fmt.Errorf("erc20 name client %w", err)
	}

	contractAddress := common.HexToAddress(contract.Address)

	instance, err := contracts.NewErc20(contractAddress, client)

	if err != nil {
		return "", fmt.Errorf("erc20 name contract %w", err)
	}

	name, err := cmn.FunctionRetrier(ctx, func() (string, error) {
		name, err := instance.Name(&bind.CallOpts{})
		return name, gw.tryWrapRetryable(ctx, "erc20 name retry", err)
	})

	if err != nil {
		return "", fmt.Errorf("erc20 name %w", err)
	}

	return name, nil
}

func (gw *gateway) GetERC20Symbol(ctx context.Context, contract *entities.Contract) (string, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return "", fmt.Errorf("erc20 symbol client %w", err)
	}

	contractAddress := common.HexToAddress(contract.Address)

	instance, err := contracts.NewErc20(contractAddress, client)

	if err != nil {
		return "", fmt.Errorf("erc20 symbol contract %w", err)
	}

	symbol, err := cmn.FunctionRetrier(ctx, func() (string, error) {
		symbol, err := instance.Symbol(&bind.CallOpts{})
		return symbol, gw.tryWrapRetryable(ctx, "erc20 symbol retry", err)
	})

	if err != nil {
		return "", fmt.Errorf("erc20 symbol %w", err)
	}

	return symbol, nil
}

func (gw *gateway) GetERC20Decimals(ctx context.Context, contract *entities.Contract) (uint8, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return 0, fmt.Errorf("erc20 decimals client %w", err)
	}

	contractAddress := common.HexToAddress(contract.Address)

	instance, err := contracts.NewErc20(contractAddress, client)

	if err != nil {
		return 0, fmt.Errorf("erc20 decimals contract %w", err)
	}

	decimals, err := cmn.FunctionRetrier(ctx, func() (uint8, error) {
		decimals, err := instance.Decimals(&bind.CallOpts{})
		return decimals, gw.tryWrapRetryable(ctx, "erc20 decimals retry", err)
	})

	if err != nil {
		return 0, fmt.Errorf("erc20 decimals %w", err)
	}

	return decimals, nil
}
