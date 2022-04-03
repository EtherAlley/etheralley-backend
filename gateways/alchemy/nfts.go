package alchemy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

type responseJson struct {
	OwnedNFTs []struct {
		Contract struct {
			Address string `json:"address"`
		} `json:"contract"`
		Id struct {
			TokenId       string `json:"tokenId"`
			TokenMetadata struct {
				TokenType string `json:"tokenType"`
			} `json:"tokenMetadata"`
		} `json:"id"`
		Metadata struct {
			Name        string                    `json:"name"`
			Description string                    `json:"description"`
			Image       string                    `json:"image"`
			Attributes  *[]map[string]interface{} `json:"attributes"`
		}
	} `json:"ownedNfts"`
}

// TODO: Polygon is also supported if we want to fetch from both in the future
func (gw *gateway) GetNonFungibleTokens(ctx context.Context, address string) *[]entities.NonFungibleToken {
	nfts := []entities.NonFungibleToken{}

	url := fmt.Sprintf("%v/getNFTs?owner=%v", gw.settings.EthereumURI(), address)

	resp, err := common.FunctionRetrier[*http.Response](ctx, gw.logger, gw.httpClient.Do, ctx, "GET", url, &common.HttpOptions{})

	if err != nil {
		gw.logger.Errf(ctx, err, "err fetching nfts from alchemy for %v, err: ", address)
		return &nfts
	}

	respJson := &responseJson{}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(respJson)

	// TODO: Taking the first 12 for now
	cutoff := len(respJson.OwnedNFTs)
	if cutoff > 12 {
		cutoff = 13
	}
	for _, nftJson := range respJson.OwnedNFTs[:cutoff] {
		nft := entities.NonFungibleToken{
			TokenId: nftJson.Id.TokenId,
			Balance: "1", // TODO: Doesn't appear that a balance is provided by alchemy
			Contract: &entities.Contract{
				Blockchain: common.ETHEREUM,
				Address:    nftJson.Contract.Address,
				Interface:  nftJson.Id.TokenMetadata.TokenType,
			},
			Metadata: &entities.NonFungibleMetadata{
				Name:        nftJson.Metadata.Name,
				Description: nftJson.Metadata.Description,
				Image:       gw.replaceIPFSScheme(nftJson.Metadata.Image),
				Attributes:  nftJson.Metadata.Attributes,
			},
		}
		nfts = append(nfts, nft)
	}

	return &nfts
}

func (gw *gateway) replaceIPFSScheme(url string) string {
	return strings.Replace(url, "ipfs://", gw.settings.IPFSURI(), 1)
}
