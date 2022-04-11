package alchemy

import (
	"context"
	"encoding/json"
	"fmt"
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
func (gw *gateway) GetNonFungibleTokens(ctx context.Context, address string) (*[]entities.NonFungibleToken, error) {
	nfts := []entities.NonFungibleToken{}

	url := fmt.Sprintf("%v/getNFTs?owner=%v", gw.settings.EthereumURI(), address)

	resp, err := gw.httpClient.Do(ctx, "GET", url, &common.HttpOptions{})

	if err != nil {
		return nil, fmt.Errorf("get all nfts %w", err)
	}

	respJson := &responseJson{}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(respJson)

	if err != nil {
		return nil, fmt.Errorf("decode all nfts response %w", err)
	}

	// TODO: Taking the first 12 for now
	cutoff := len(respJson.OwnedNFTs)
	if cutoff > 12 {
		cutoff = 13
	}
	for _, nftJson := range respJson.OwnedNFTs[:cutoff] {
		balance := "1"
		nft := entities.NonFungibleToken{
			TokenId: nftJson.Id.TokenId,
			Balance: &balance, // TODO: Doesn't appear that a balance is provided by alchemy for ERC1155, only ERC721...
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

	return &nfts, nil
}

type nftMetadataRespBody struct {
	Name        string                    `bson:"name" json:"name"`
	Description string                    `bson:"description" json:"description"`
	Image       string                    `bson:"image" json:"image"`
	ImageURL    string                    `bson:"image_url" json:"image_url"`
	Attributes  *[]map[string]interface{} `bson:"attributes" json:"attributes"`
	Properties  *map[string]interface{}   `bson:"properties" json:"properties"`
}

func (gw *gateway) GetNonFungibleMetadata(ctx context.Context, uri string) (*entities.NonFungibleMetadata, error) {
	uri = gw.replaceIPFSScheme(uri)

	resp, err := gw.httpClient.Do(ctx, "GET", uri, nil)

	if err != nil {
		return nil, fmt.Errorf("metadata follow url %w", err)
	}

	metadata := &nftMetadataRespBody{}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&metadata)

	if err != nil {
		return nil, fmt.Errorf("metadata decode response %w", err)
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
