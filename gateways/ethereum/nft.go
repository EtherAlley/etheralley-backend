package ethereum

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"strings"

	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (gw *Gateway) getNFTMetadataFromURI(uri string) (*gateways.NFTMetadata, error) {
	metadata := &gateways.NFTMetadata{}

	uri = strings.Replace(uri, "ipfs://", "https://ipfs.io/ipfs/", 1)

	gw.logger.Infof("http get: %v", uri)

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

func (gw *Gateway) GetERC1155NFTMetadata(contractAddress string, tokenId string) (*gateways.NFTMetadata, error) {
	address := common.HexToAddress(contractAddress)
	instance, err := contracts.NewErc1155(address, gw.client)

	if err != nil {
		return nil, err
	}

	id := new(big.Int)
	id, ok := id.SetString(tokenId, 10)

	if !ok {
		return nil, errors.New("invalid token id")
	}

	uri, err := instance.Uri(&bind.CallOpts{}, id)

	if err != nil {
		return nil, err
	}

	gw.logger.Info(uri)

	hexId := hex.EncodeToString(id.Bytes())
	uri = strings.Replace(uri, "{id}", hexId, 1)

	gw.logger.Info(uri)

	return gw.getNFTMetadataFromURI(uri)
}

func (gw *Gateway) GetERC721NFTMetadata(contractAddress string, tokenId string) (*gateways.NFTMetadata, error) {
	address := common.HexToAddress(contractAddress)
	instance, err := contracts.NewErc721(address, gw.client)

	if err != nil {
		return nil, err
	}

	id := new(big.Int)
	id, ok := id.SetString(tokenId, 10)

	if !ok {
		return nil, errors.New("invalid token id")
	}

	uri, err := instance.TokenURI(&bind.CallOpts{}, id)

	if err != nil {
		return nil, err
	}

	return gw.getNFTMetadataFromURI(uri)
}

func (gw *Gateway) VerifyERC1155Owner(contractAddress string, address string, tokenId string) (bool, error) {
	return false, nil
}

func (gw *Gateway) VerifyERC721Owner(contractAddress string, address string, tokenId string) (bool, error) {
	return false, nil
}
