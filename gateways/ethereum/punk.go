package ethereum

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (gw *gateway) GetPunkBalance(ctx context.Context, address string, contract *entities.Contract, tokenId string) (string, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return "", fmt.Errorf("punk balanceOf client %w", err)
	}

	contractAddress := common.HexToAddress(contract.Address)
	adr := common.HexToAddress(address)
	id := new(big.Int)
	id, ok := id.SetString(tokenId, 10)

	if !ok {
		return "", errors.New("punk balanceOf invalid token id")
	}

	instance, err := contracts.NewCryptoPunks(contractAddress, client)

	if err != nil {
		return "", fmt.Errorf("punk balanceOf contract %w", err)
	}

	owner, err := cmn.FunctionRetrier(ctx, func() (common.Address, error) {
		owner, err := instance.PunkIndexToAddress(&bind.CallOpts{}, id)
		return owner, gw.tryWrapRetryable(ctx, "punk balanceOf retry", err)
	})

	// treat a bad contract address as a zero balance
	if errors.Is(err, bind.ErrNoCode) {
		return "0", nil
	}

	if err != nil {
		return "", fmt.Errorf("punk balanceOf %w", err)
	}

	if adr.Hex() == owner.Hex() {
		return "1", nil
	} else {
		return "0", nil
	}
}
