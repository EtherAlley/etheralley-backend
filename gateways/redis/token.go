package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/etheralley/etheralley-core-api/entities"
)

const TokenNamespace = "fungible_token"

func (g *gateway) GetFungibleMetadata(ctx context.Context, contract *entities.Contract) (*entities.FungibleMetadata, error) {

	metadataString, err := g.client.Get(ctx, getFullKey(TokenNamespace, contract.Address, contract.Blockchain)).Result()

	if err != nil {
		return nil, fmt.Errorf("get fungible metadata %w", err)
	}

	metadataJson := &fungibleMetadataJson{}
	err = json.Unmarshal([]byte(metadataString), metadataJson)

	if err != nil {
		return nil, fmt.Errorf("decode fungible metadata %w", err)
	}

	metadata := fromFungibleMetadataJson(metadataJson)

	return metadata, nil
}

func (g *gateway) SaveFungibleMetadata(ctx context.Context, contract *entities.Contract, metadata *entities.FungibleMetadata) error {
	metadataJson := toFungibleMetadataJson(metadata)
	bytes, err := json.Marshal(metadataJson)

	if err != nil {
		return fmt.Errorf("encode fungible metadata %w", err)
	}

	_, err = g.client.Set(ctx, getFullKey(TokenNamespace, contract.Address, contract.Blockchain), bytes, time.Hour*24).Result()

	if err != nil {
		return fmt.Errorf("save fungible metadata %w", err)
	}

	return nil
}
