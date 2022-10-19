package offchain

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

type alchemyGetAllNFTsResponseJson struct {
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
		Metadata json.RawMessage `json:"metadata"` // This can be either string or struct. We need to parse one by one
		Error    string          `json:"error"`
	} `json:"ownedNfts"`
}

type responseMetadataJson struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Image       string                    `json:"image"`
	ImageURL    string                    `json:"image_url"`
	Attributes  *[]map[string]interface{} `json:"attributes"`
}

// See https://docs.alchemy.com/alchemy/enhanced-apis/nft-api/getnfts
// TODO: Polygon is also supported if we want to fetch from both in the future
func (gw *gateway) GetNonFungibleTokens(ctx context.Context, address string) (*[]entities.NonFungibleToken, error) {
	url := fmt.Sprintf("%v/getNFTs?owner=%v&filters[]=SPAM", gw.settings.AlchemyEthereumURI(), address)

	resp, err := gw.httpClient.Do(ctx, "GET", url, nil, &common.HttpOptions{})

	if err != nil {
		return nil, fmt.Errorf("get all nfts %w", err)
	}

	respJson := &alchemyGetAllNFTsResponseJson{}
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

		// skip anything that is not the standard metadata json structur
		var metadata responseMetadataJson
		if err := json.Unmarshal(nftJson.Metadata, &metadata); err != nil {
			continue
		}

		image := ""
		if metadata.Image != "" {
			image = metadata.Image
		} else if metadata.ImageURL != "" {
			image = metadata.ImageURL
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

		// we handle ens nfts specially because we need special logic to fetch their metadata
		tokenType := nftJson.Id.TokenMetadata.TokenType
		if nftJson.Contract.Address == common.ENS_BASE_REGISTRAR_ADDRESS {
			tokenType = common.ENS_REGISTRAR
		}

		nft := entities.NonFungibleToken{
			TokenId: tokenId,
			Balance: &balance,
			Contract: &entities.Contract{
				Blockchain: common.ETHEREUM,
				Address:    nftJson.Contract.Address,
				Interface:  tokenType,
			},
			Metadata: &entities.NonFungibleMetadata{
				Name:        metadata.Name,
				Description: metadata.Description,
				Image:       image,
				Attributes:  metadata.Attributes,
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

	resp, err := gw.httpClient.Do(ctx, "GET", uri, nil, nil)

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
