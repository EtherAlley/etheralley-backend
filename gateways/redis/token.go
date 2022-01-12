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

	metadata := &entities.FungibleMetadata{}
	err = json.Unmarshal([]byte(metadataString), &metadata)

	if err != nil {
		return nil, err
	}

	return metadata, err
}

func (g *Gateway) SaveFungibleMetadata(ctx context.Context, contract *entities.Contract, metadata *entities.FungibleMetadata) error {
	metadataBytes, err := json.Marshal(metadata)

	if err != nil {
		return err
	}

	_, err = g.client.Set(ctx, getFullKey(TokenNamespace, contract.Address, contract.Blockchain), string(metadataBytes), time.Hour*24).Result()

	return err
}
