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
		Error string `json:"error"`
	} `json:"ownedNfts"`
}

// See https://docs.alchemy.com/alchemy/enhanced-apis/nft-api/getnfts
// TODO: Polygon is also supported if we want to fetch from both in the future
func (gw *gateway) GetNonFungibleTokens(ctx context.Context, address string) (*[]entities.NonFungibleToken, error) {
	url := fmt.Sprintf("%v/getNFTs?owner=%v&filters[]=SPAM", gw.settings.EthereumURI(), address)

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

	nfts := []entities.NonFungibleToken{}
	for _, nftJson := range respJson.OwnedNFTs {
		// skip nfts that have an error
		if nftJson.Error != "" {
			continue
		}

		image := ""
		if nftJson.Metadata.Image != "" {
			image = nftJson.Metadata.Image
		} else if nftJson.Metadata.ImageURL != "" {
			image = nftJson.Metadata.ImageURL
		} else { // skip nfts that don't have an image url
			continue
		}
		image = gw.replaceIPFSScheme(image)

		// It appears the token ids are provided in hexidecimal format
		tokenId := nftJson.Id.TokenId
		if strings.Contains(tokenId, "0x") {
			num := new(big.Int)
			num, ok := num.SetString(strings.Split(tokenId, "0x")[1], 16)
			if !ok {
				return nil, fmt.Errorf("parsing token id %v", tokenId)
			}
			tokenId = num.String()
		}

		// Doesn't appear that a balance is provided by alchemy for ERC1155, only ERC721... Hardcoding balance of 1 for now.
		balance := "1"

		nft := entities.NonFungibleToken{
			TokenId: tokenId,
			Balance: &balance,
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

	// TODO: Taking the first 12 for now
	// Reminder, making this number bigger than 12 has implications for the max badge limit
	cutoff := len(nfts)
	if cutoff > 12 {
		cutoff = 13
	}
	trimmedNFTs := nfts[:cutoff]

	return &trimmedNFTs, nil
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
