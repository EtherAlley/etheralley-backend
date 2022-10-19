package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

const NFTNamespace = "non_fungible_token"

func (g *gateway) GetNonFungibleMetadata(ctx context.Context, contract *entities.Contract, tokenId string) (*entities.NonFungibleMetadata, error) {

	metadataString, err := g.client.Get(ctx, getFullKey(NFTNamespace, contract.Address, tokenId, contract.Blockchain)).Result()

	if err != nil {
		return nil, fmt.Errorf("get nft metadata %w", err)
	}

	metadataJson := &nonFungibleMetadataJson{}
	err = json.Unmarshal([]byte(metadataString), metadataJson)

	if err != nil {
		return nil, fmt.Errorf("decode nft metadata %w", err)
	}

	metadata := fromNonFungibleMetadataJson(metadataJson)

	return metadata, nil
}

func (g *gateway) SaveNonFungibleMetadata(ctx context.Context, contract *entities.Contract, tokenId string, metadata *entities.NonFungibleMetadata) error {
	metadataJson := toNonFungibleMetadataJson(metadata)
	bytes, err := json.Marshal(metadataJson)

	if err != nil {
		return fmt.Errorf("encode nft metadata %w", err)
	}

	_, err = g.client.Set(ctx, getFullKey(NFTNamespace, contract.Address, tokenId, contract.Blockchain), bytes, time.Hour).Result()

	if err != nil {
		return fmt.Errorf("save nft metadata  %w", err)
	}

	return nil
}
