package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/etheralley/etheralley-core-api/entities"
)

const NFTNamespace = "non_fungible_token"

func (g *Gateway) GetNonFungibleMetadata(ctx context.Context, contract *entities.Contract, tokenId string) (*entities.NonFungibleMetadata, error) {

	metadataString, err := g.client.Get(ctx, getFullKey(NFTNamespace, contract.Address, tokenId, contract.Blockchain)).Result()

	if err != nil {
		return nil, err
	}

	metadataJson := &nonFungibleMetadataJson{}
	err = json.Unmarshal([]byte(metadataString), &metadataJson)

	if err != nil {
		return nil, err
	}

	metadata := fromNonFungibleMetadataJson(metadataJson)

	return metadata, nil
}

func (g *Gateway) SaveNonFungibleMetadata(ctx context.Context, contract *entities.Contract, tokenId string, metadata *entities.NonFungibleMetadata) error {
	metadataJson := toNonFungibleMetadataJson(metadata)
	bytes, err := json.Marshal(metadataJson)

	if err != nil {
		return err
	}

	_, err = g.client.Set(ctx, getFullKey(NFTNamespace, contract.Address, tokenId, contract.Blockchain), bytes, time.Hour*24).Result()

	return err
}

func fromNonFungibleMetadataJson(metadata *nonFungibleMetadataJson) *entities.NonFungibleMetadata {
	return &entities.NonFungibleMetadata{
		Name:        metadata.Name,
		Description: metadata.Description,
		Image:       metadata.Image,
		Attributes:  metadata.Attributes,
	}
}

func toNonFungibleMetadataJson(metadata *entities.NonFungibleMetadata) *nonFungibleMetadataJson {
	return &nonFungibleMetadataJson{
		Name:        metadata.Name,
		Description: metadata.Description,
		Image:       metadata.Image,
		Attributes:  metadata.Attributes,
	}
}
