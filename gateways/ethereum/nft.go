package ethereum

import (
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

func (gw *Gateway) GetNonFungibleMetadata(contract *entities.Contract, tokenId string) (*entities.NonFungibleMetadata, error) {
	client, err := gw.getClient(contract.Blockchain)

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
		uri, err = gw.getErc721URI(client, address, id)
	case cmn.ERC1155:
		uri, err = gw.getErc1155URI(client, address, id)
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

	return gw.getNFTMetadataFromURI(uri)
}

func (gw *Gateway) GetNonFungibleBalance(address string, contract *entities.Contract, tokenId string) (string, error) {
	client, err := gw.getClient(contract.Blockchain)

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
		return gw.getErc1155Balance(client, contractAddress, adr, id)
	case cmn.ERC721, cmn.ENS_REGISTRAR:
		return gw.getErc721Balance(client, contractAddress, adr, id)
	default:
		return "", errors.New("invalida schema name")
	}
}

func (gw *Gateway) getErc1155Balance(client *ethclient.Client, contractAddress common.Address, address common.Address, tokenId *big.Int) (string, error) {
	instance, err := contracts.NewErc1155(contractAddress, client)

	if err != nil {
		return "", err
	}

	balance, err := instance.BalanceOf(&bind.CallOpts{}, address, tokenId)

	if err != nil {
		return "", err
	}

	return balance.String(), err
}

func (gw *Gateway) getErc721Balance(client *ethclient.Client, contractAddress common.Address, address common.Address, tokenId *big.Int) (string, error) {
	instance, err := contracts.NewErc721(contractAddress, client)

	if err != nil {
		return "", err
	}

	owner, err := instance.OwnerOf(&bind.CallOpts{}, tokenId)

	if err != nil {
		return "", err
	}

	if address.Hex() == owner.Hex() {
		return "1", nil
	} else {
		return "0", nil
	}
}

func (gw *Gateway) getErc1155URI(client *ethclient.Client, address common.Address, id *big.Int) (string, error) {
	instance, err := contracts.NewErc1155(address, client)

	if err != nil {
		return "", err
	}

	uri, err := instance.Uri(&bind.CallOpts{}, id)

	if err != nil {
		return "", err
	}

	hexId := hex.EncodeToString(id.Bytes())
	uri = strings.Replace(uri, "{id}", hexId, 1)

	return uri, nil
}

func (gw *Gateway) getErc721URI(client *ethclient.Client, address common.Address, id *big.Int) (string, error) {
	instance, err := contracts.NewErc721(address, client)

	if err != nil {
		return "", err
	}

	return instance.TokenURI(&bind.CallOpts{}, id)
}

func (gw *Gateway) getENSURI(client *ethclient.Client, address string, id string) string {
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

func (gw *Gateway) getNFTMetadataFromURI(uri string) (*entities.NonFungibleMetadata, error) {
	uri = gw.replaceIPFSScheme(uri)

	resp, err := gw.http.Do("GET", uri, nil)

	if err != nil {
		gw.logger.Errf(err, "nft metadata url follow http err: ")
		return nil, errors.New("could not fetch metadata url")
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

func (gw *Gateway) replaceIPFSScheme(url string) string {
	return strings.Replace(url, "ipfs://", gw.settings.IPFSURI(), 1)
}
