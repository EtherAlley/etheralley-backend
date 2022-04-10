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
	"github.com/ethereum/go-ethereum/ethclient"
)

func (gw *gateway) GetNonFungibleURI(ctx context.Context, contract *entities.Contract, tokenId string) (string, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return "", fmt.Errorf("nft metadata client %w", err)
	}

	address := common.HexToAddress(contract.Address)
	id := new(big.Int)
	id, ok := id.SetString(tokenId, 10)

	if !ok {
		return "", errors.New("invalid token id")
	}

	switch contract.Interface {
	case cmn.ERC721:
		return gw.getErc721URI(ctx, client, address, id)
	case cmn.ERC1155:
		return gw.getErc1155URI(ctx, client, address, id)
	case cmn.ENS_REGISTRAR:
		return gw.getENSURI(contract.Address, tokenId), nil
	default:
		return "", errors.New("invalida schema name")
	}
}

func (gw *gateway) GetNonFungibleBalance(ctx context.Context, address string, contract *entities.Contract, tokenId string) (string, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return "", fmt.Errorf("balance client %w", err)
	}

	contractAddress := common.HexToAddress(contract.Address)
	adr := common.HexToAddress(address)
	id := new(big.Int)
	id, ok := id.SetString(tokenId, 10)

	if !ok {
		return "", errors.New("invalid token id")
	}

	switch contract.Interface {
	case cmn.ERC1155:
		return gw.getErc1155Balance(ctx, client, contractAddress, adr, id)
	case cmn.ERC721, cmn.ENS_REGISTRAR:
		return gw.getErc721Balance(ctx, client, contractAddress, adr, id)
	default:
		return "", errors.New("invalida schema name")
	}
}

func (gw *gateway) getErc1155Balance(ctx context.Context, client *ethclient.Client, contractAddress common.Address, address common.Address, tokenId *big.Int) (string, error) {
	instance, err := contracts.NewErc1155(contractAddress, client)

	if err != nil {
		return "", fmt.Errorf("erc1155 balanceOf contract %w", err)
	}

	balance, err := cmn.FunctionRetrier(ctx, func() (*big.Int, error) {
		balance, err := instance.BalanceOf(&bind.CallOpts{}, address, tokenId)
		return balance, tryWrapRetryable("erc1155 balanceOf retry", err)
	})

	if err != nil {
		return "", fmt.Errorf("erc1155 balanceOf %w", err)
	}

	return balance.String(), err
}

func (gw *gateway) getErc721Balance(ctx context.Context, client *ethclient.Client, contractAddress common.Address, address common.Address, tokenId *big.Int) (string, error) {
	instance, err := contracts.NewErc721(contractAddress, client)

	if err != nil {
		return "", fmt.Errorf("erc721 balanceOf contract %w", err)
	}

	owner, err := cmn.FunctionRetrier(ctx, func() (common.Address, error) {
		owner, err := instance.OwnerOf(&bind.CallOpts{}, tokenId)
		return owner, tryWrapRetryable("erc721 balanceOf retry", err)
	})

	if err != nil {
		return "", fmt.Errorf("erc721 balanceOf %w", err)
	}

	if address.Hex() == owner.Hex() {
		return "1", nil
	} else {
		return "0", nil
	}
}

func (gw *gateway) getErc1155URI(ctx context.Context, client *ethclient.Client, address common.Address, id *big.Int) (string, error) {
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

func (gw *gateway) getErc721URI(ctx context.Context, client *ethclient.Client, address common.Address, id *big.Int) (string, error) {
	instance, err := contracts.NewErc721(address, client)

	if err != nil {
		return "", fmt.Errorf("erc721 uri contract %w", err)
	}

	return cmn.FunctionRetrier(ctx, func() (string, error) {
		uri, err := instance.TokenURI(&bind.CallOpts{}, id)
		return uri, tryWrapRetryable("erc721 uri retry", err)
	})
}

func (gw *gateway) getENSURI(address string, id string) string {
	return fmt.Sprintf("%v/%v/%v", gw.settings.ENSMetadataURI(), address, id)
}
