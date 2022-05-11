package offchain

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
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
			ImageURL    string                    `json:"image_url"`
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

		image := ""
		if nftJson.Metadata.Image != "" {
			image = nftJson.Metadata.Image
		} else if nftJson.Metadata.ImageURL != "" {
			image = nftJson.Metadata.ImageURL
		}
		image = gw.replaceIPFSScheme(image)

		// Appears the token ids are provided in hexidecimal format
		tokenId := nftJson.Id.TokenId
		if strings.Contains(tokenId, "0x") {
			num := new(big.Int)
			num, ok := num.SetString(strings.Split(tokenId, "0x")[1], 16)
			if !ok {
				return nil, fmt.Errorf("parsing token id %v", tokenId)
			}
			tokenId = num.String()
		}

		nft := entities.NonFungibleToken{
			TokenId: tokenId,
			Balance: &balance, // Doesn't appear that a balance is provided by alchemy for ERC1155, only ERC721...
			Contract: &entities.Contract{
				Blockchain: common.ETHEREUM,
				Address:    nftJson.Contract.Address,
				Interface:  nftJson.Id.TokenMetadata.TokenType,
			},
			Metadata: &entities.NonFungibleMetadata{
				Name:        nftJson.Metadata.Name,
				Description: nftJson.Metadata.Description,
				Image:       image,
				Attributes:  nftJson.Metadata.Attributes,
			},
		}
		nfts = append(nfts, nft)
	}

	return &nfts, nil
}

type nftMetadataRespBody struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Image       string                    `json:"image"`
	ImageURL    string                    `json:"image_url"`
	Attributes  *[]map[string]interface{} `json:"attributes"`
	Properties  *map[string]interface{}   `json:"properties"`
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
