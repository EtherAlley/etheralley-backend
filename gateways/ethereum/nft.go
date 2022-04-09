package ethereum

import (
	"context"
	"encoding/hex"
	"encoding/json"
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

func (gw *gateway) GetNonFungibleMetadata(ctx context.Context, contract *entities.Contract, tokenId string) (*entities.NonFungibleMetadata, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(contract.Address)
	id := new(big.Int)
	id, ok := id.SetString(tokenId, 10)

	if !ok {
		return nil, errors.New("invalid token id")
	}

	var uri string
	switch contract.Interface {
	case cmn.ERC721:
		uri, err = gw.getErc721URI(ctx, client, address, id)
	case cmn.ERC1155:
		uri, err = gw.getErc1155URI(ctx, client, address, id)
	case cmn.ENS_REGISTRAR:
		uri = gw.getENSURI(client, contract.Address, tokenId)
		err = nil
	default:
		uri = ""
		err = errors.New("invalida schema name")
	}

	if err != nil {
		return nil, err
	}

	return gw.getNFTMetadataFromURI(ctx, uri)
}

func (gw *gateway) GetNonFungibleBalance(ctx context.Context, address string, contract *entities.Contract, tokenId string) (string, error) {
	client, err := gw.getClient(ctx, contract.Blockchain)

	if err != nil {
		return "", err
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
		return "", err
	}

	balance, err := cmn.FunctionRetrier(ctx, func() (*big.Int, error) {
		balance, err := instance.BalanceOf(&bind.CallOpts{}, address, tokenId)
		return balance, tryWrapRetryable("get nft erc1155 balance", err)
	})

	if err != nil {
		return "", err
	}

	return balance.String(), err
}

func (gw *gateway) getErc721Balance(ctx context.Context, client *ethclient.Client, contractAddress common.Address, address common.Address, tokenId *big.Int) (string, error) {
	instance, err := contracts.NewErc721(contractAddress, client)

	if err != nil {
		return "", err
	}

	owner, err := cmn.FunctionRetrier(ctx, func() (common.Address, error) {
		owner, err := instance.OwnerOf(&bind.CallOpts{}, tokenId)
		return owner, tryWrapRetryable("get nft erc721 balance", err)
	})

	if err != nil {
		return "", err
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
		return "", err
	}

	uri, err := cmn.FunctionRetrier(ctx, func() (string, error) {
		uri, err := instance.Uri(&bind.CallOpts{}, id)
		return uri, tryWrapRetryable("get nft 1155 uri", err)
	})

	if err != nil {
		return "", err
	}

	// See https://eips.ethereum.org/EIPS/eip-1155: token ids are passed in hexidecimal form
	hexId := hex.EncodeToString(id.Bytes())
	uri = strings.Replace(uri, "{id}", hexId, 1)

	return uri, nil
}

func (gw *gateway) getErc721URI(ctx context.Context, client *ethclient.Client, address common.Address, id *big.Int) (string, error) {
	instance, err := contracts.NewErc721(address, client)

	if err != nil {
		return "", err
	}

	return cmn.FunctionRetrier(ctx, func() (string, error) {
		uri, err := instance.TokenURI(&bind.CallOpts{}, id)
		return uri, tryWrapRetryable("get nft erc721 uri", err)
	})
}

func (gw *gateway) getENSURI(client *ethclient.Client, address string, id string) string {
	return fmt.Sprintf("%v/%v/%v", gw.settings.ENSMetadataURI(), address, id)
}

type NFTMetadataRespBody struct {
	Name        string                    `bson:"name" json:"name"`
	Description string                    `bson:"description" json:"description"`
	Image       string                    `bson:"image" json:"image"`
	ImageURL    string                    `bson:"image_url" json:"image_url"`
	Attributes  *[]map[string]interface{} `bson:"attributes" json:"attributes"`
	Properties  *map[string]interface{}   `bson:"properties" json:"properties"`
}

func (gw *gateway) getNFTMetadataFromURI(ctx context.Context, uri string) (*entities.NonFungibleMetadata, error) {
	uri = gw.replaceIPFSScheme(uri)

	resp, err := gw.http.Do(ctx, "GET", uri, nil)

	if err != nil {
		gw.logger.Errf(ctx, err, "nft metadata url follow http err: ")
		return nil, fmt.Errorf("could not fetch metadata url %w", err)
	}

	metadata := &NFTMetadataRespBody{}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&metadata)

	if err != nil {
		return nil, err
	}

	image := ""
	if metadata.Image != "" {
		image = metadata.Image
	} else if metadata.ImageURL != "" {
		image = metadata.ImageURL
	}
	image = gw.replaceIPFSScheme(image)

	return &entities.NonFungibleMetadata{
		Name:        metadata.Name,
		Description: metadata.Description,
		Image:       image,
		Attributes:  metadata.Attributes,
	}, nil
}

func (gw *gateway) replaceIPFSScheme(url string) string {
	return strings.Replace(url, "ipfs://", gw.settings.IPFSURI(), 1)
}
