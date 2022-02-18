package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/etheralley/etheralley-core-api/entities"
)

const TokenNamespace = "fungible_token"

func (g *Gateway) GetFungibleMetadata(ctx context.Context, contract *entities.Contract) (*entities.FungibleMetadata, error) {

	metadataString, err := g.client.Get(ctx, getFullKey(TokenNamespace, contract.Address, contract.Blockchain)).Result()

	if err != nil {
		return nil, err
	}

	metadataJson := &fungibleMetadataJson{}
	err = json.Unmarshal([]byte(metadataString), &metadataJson)

	if err != nil {
		return nil, err
	}

	metadata := fromFungibleMetadataJson(metadataJson)

	return metadata, nil
}

func (g *Gateway) SaveFungibleMetadata(ctx context.Context, contract *entities.Contract, metadata *entities.FungibleMetadata) error {
	metadataJson := toFungibleMetadataJson(metadata)
	bytes, err := json.Marshal(metadataJson)

	if err != nil {
		return err
	}

	_, err = g.client.Set(ctx, getFullKey(TokenNamespace, contract.Address, contract.Blockchain), bytes, time.Hour*24).Result()

	return err
}

func fromFungibleMetadataJson(metadata *fungibleMetadataJson) *entities.FungibleMetadata {
	return &entities.FungibleMetadata{
		Name:     metadata.Name,
		Symbol:   metadata.Symbol,
		Decimals: metadata.Decimals,
	}
}

func toFungibleMetadataJson(metadata *entities.FungibleMetadata) *fungibleMetadataJson {
	return &fungibleMetadataJson{
		Name:     metadata.Name,
		Symbol:   metadata.Symbol,
		Decimals: metadata.Decimals,
	}
}
