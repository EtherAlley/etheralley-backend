package thegraph

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

type ERC721Token struct {
	Id       string
	Contract struct {
		Id string
	}
	Identifier string
	Owner      struct {
		Id string
	}
	Uri string
}

type ERC721Query = struct {
	Erc721Tokens []ERC721Token `graphql:"erc721Tokens(first: 12, where: { owner: $owner, uri_contains: \"://\" })"`
}

func (gw *Gateway) GetNonFungibleTokens(ctx context.Context, address string) *[]entities.NonFungibleToken {
	nfts := []entities.NonFungibleToken{}

	url, err := gw.GetSubgraphUrl(common.ETHEREUM, common.ERC721)

	if err != nil {
		gw.logger.Errf(err, "error building subgraph url for address: %v", address)
		return &nfts
	}

	query := &ERC721Query{}
	variables := map[string]interface{}{
		"owner": strings.ToLower(address), // not sure why this is needed. Isnt this against some kind of standard that we have to do this?
	}
	err = gw.graphClient.Query(ctx, url, query, variables)

	if err != nil {
		gw.logger.Errf(err, "error calling subgraph for address: %v", address)
		return &nfts
	}

	var wg sync.WaitGroup

	tokens := make([]*entities.NonFungibleToken, len(query.Erc721Tokens))

	for i, token := range query.Erc721Tokens {
		wg.Add(1)

		go func(i int, t ERC721Token) {
			defer wg.Done()

			metadata, err := gw.getNFTMetadataFromURI(t.Uri)

			if err != nil {
				gw.logger.Errf(err, "err fetching nft metadata: contract address %v token id %v user address %v", t.Contract.Id, t.Identifier, address)
				return
			}

			tokens[i] = &entities.NonFungibleToken{
				Contract: &entities.Contract{
					Blockchain: common.ETHEREUM,
					Address:    t.Contract.Id,
					Interface:  common.ERC721,
				},
				TokenId:  t.Identifier,
				Balance:  "1",
				Metadata: metadata,
			}
		}(i, token)
	}

	wg.Wait()

	for _, token := range tokens {
		if token != nil {
			nfts = append(nfts, *token)
		}
	}

	return &nfts
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

	resp, err := gw.httpClient.Do("GET", uri, nil)

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
	return strings.Replace(url, "ipfs://", gw.settings.IPFSURI, 1)
}
