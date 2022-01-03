package ethereum

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"strings"

	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (gw *Gateway) GetNFTMetadata(location *entities.NFTLocation) (*entities.NFTMetadata, error) {
	address := common.HexToAddress(location.ContractAddress)
	id := new(big.Int)
	id, ok := id.SetString(location.TokenId, 10)

	if !ok {
		return nil, errors.New("invalid token id")
	}

	var uri string
	var err error
	switch location.SchemaName {
	case "ERC721":
		uri, err = gw.getErc721URI(address, id)
	case "ERC1155":
		uri, err = gw.getErc1155URI(address, id)
	default:
		uri = ""
		err = errors.New("invalida schema name")
	}

	if err != nil {
		return nil, err
	}

	return gw.getNFTMetadataFromURI(uri)
}

func (gw *Gateway) VerifyOwner(address string, location *entities.NFTLocation) (bool, error) {
	contractAddress := common.HexToAddress(location.ContractAddress)
	adr := common.HexToAddress(address)
	id := new(big.Int)
	id, ok := id.SetString(location.TokenId, 10)

	if !ok {
		return false, errors.New("invalid token id")
	}

	switch location.SchemaName {
	case "ERC1155":
		return gw.verifyErc1155Owner(contractAddress, adr, id)
	case "ERC721":
		return gw.verifyErc721Owner(contractAddress, adr, id)
	default:
		return false, errors.New("invalida schema name")
	}
}

func (gw *Gateway) verifyErc1155Owner(contractAddress common.Address, address common.Address, tokenId *big.Int) (bool, error) {
	instance, err := contracts.NewErc1155(contractAddress, gw.client)

	if err != nil {
		return false, err
	}

	balance, err := instance.BalanceOf(&bind.CallOpts{}, address, tokenId)

	return balance.Cmp(big.NewInt(0)) == 1, err
}

func (gw *Gateway) verifyErc721Owner(contractAddress common.Address, address common.Address, tokenId *big.Int) (bool, error) {
	instance, err := contracts.NewErc721(contractAddress, gw.client)

	if err != nil {
		return false, err
	}

	owner, err := instance.OwnerOf(&bind.CallOpts{}, tokenId)

	return address.Hex() == owner.Hex(), err
}

func (gw *Gateway) getErc1155URI(address common.Address, id *big.Int) (string, error) {
	instance, err := contracts.NewErc1155(address, gw.client)

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

func (gw *Gateway) getErc721URI(address common.Address, id *big.Int) (string, error) {
	instance, err := contracts.NewErc721(address, gw.client)

	if err != nil {
		return "", err
	}

	return instance.TokenURI(&bind.CallOpts{}, id)
}

func (gw *Gateway) getNFTMetadataFromURI(uri string) (*entities.NFTMetadata, error) {
	metadata := &entities.NFTMetadata{}

	uri = strings.Replace(uri, "ipfs://", "https://ipfs.io/ipfs/", 1)

	gw.logger.Debugf("ethereum gateway: http get: %v", uri)

	resp, err := http.Get(uri)

	if err != nil || resp.StatusCode != 200 {
		return nil, errors.New("could not fetch uri")
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&metadata)

	if err != nil {
		return nil, err
	}

	return metadata, nil
}
