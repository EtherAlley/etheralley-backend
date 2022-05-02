package ethereum

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	cmn "github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (gw *gateway) GetERC1155Balance(ctx context.Context, address string, contract *entities.Contract, tokenId string) (string, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return "", fmt.Errorf("erc1155 balance balance client %w", err)
	}

	contractAddress := common.HexToAddress(contract.Address)
	adr := common.HexToAddress(address)
	id := new(big.Int)
	id, ok := id.SetString(tokenId, 10)

	if !ok {
		return "", errors.New("erc1155 balance invalid token id")
	}

	instance, err := contracts.NewErc1155(contractAddress, client)

	if err != nil {
		return "", fmt.Errorf("erc1155 balance contract %w", err)
	}

	balance, err := cmn.FunctionRetrier(ctx, func() (*big.Int, error) {
		balance, err := instance.BalanceOf(&bind.CallOpts{}, adr, id)
		return balance, tryWrapRetryable("erc1155 balance retry", err)
	})

	// treat a bad contract address as a zero balance
	if errors.Is(err, bind.ErrNoCode) {
		return "0", nil
	}

	if err != nil {
		return "", fmt.Errorf("erc1155 balance %w", err)
	}

	return balance.String(), err
}

func (gw *gateway) GetERC721Balance(ctx context.Context, address string, contract *entities.Contract, tokenId string) (string, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return "", fmt.Errorf("erc721 balance client %w", err)
	}

	contractAddress := common.HexToAddress(contract.Address)
	adr := common.HexToAddress(address)
	id := new(big.Int)
	id, ok := id.SetString(tokenId, 10)

	if !ok {
		return "", errors.New("erc721 balance invalid token id")
	}

	instance, err := contracts.NewErc721(contractAddress, client)

	if err != nil {
		return "", fmt.Errorf("erc721 balance contract %w", err)
	}

	owner, err := cmn.FunctionRetrier(ctx, func() (common.Address, error) {
		owner, err := instance.OwnerOf(&bind.CallOpts{}, id)
		return owner, tryWrapRetryable("erc721 balance retry", err)
	})

	// treat a bad contract address as a zero balance
	if errors.Is(err, bind.ErrNoCode) {
		return "0", nil
	}

	if err != nil {
		return "", fmt.Errorf("erc721 balanceOf %w", err)
	}

	if adr.Hex() == owner.Hex() {
		return "1", nil
	} else {
		return "0", nil
	}
}

func (gw *gateway) GetERC1155URI(ctx context.Context, contract *entities.Contract, tokenId string) (string, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return "", fmt.Errorf("erc1155 uri client %w", err)
	}

	address := common.HexToAddress(contract.Address)
	id := new(big.Int)
	id, ok := id.SetString(tokenId, 10)

	if !ok {
		return "", errors.New("erc1155 uri invalid token id")
	}

	instance, err := contracts.NewErc1155(address, client)

	if err != nil {
		return "", fmt.Errorf("erc1155 uri contract %w", err)
	}

	uri, err := cmn.FunctionRetrier(ctx, func() (string, error) {
		uri, err := instance.Uri(&bind.CallOpts{}, id)
		return uri, tryWrapRetryable("erc1155 uri retry", err)
	})

	if err != nil {
		return "", fmt.Errorf("erc1155 uri %w", err)
	}

	// See https://eips.ethereum.org/EIPS/eip-1155: token ids are passed in hexidecimal form
	hexId := hex.EncodeToString(id.Bytes())
	uri = strings.Replace(uri, "{id}", hexId, 1)

	return uri, nil
}

func (gw *gateway) GetERC721URI(ctx context.Context, contract *entities.Contract, tokenId string) (string, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return "", fmt.Errorf("erc721 uri client %w", err)
	}

	address := common.HexToAddress(contract.Address)
	id := new(big.Int)
	id, ok := id.SetString(tokenId, 10)

	if !ok {
		return "", errors.New("erc721 uri invalid token id")
	}

	instance, err := contracts.NewErc721(address, client)

	if err != nil {
		return "", fmt.Errorf("erc721 uri contract %w", err)
	}

	return cmn.FunctionRetrier(ctx, func() (string, error) {
		uri, err := instance.TokenURI(&bind.CallOpts{}, id)
		return uri, tryWrapRetryable("erc721 uri retry", err)
	})
}
