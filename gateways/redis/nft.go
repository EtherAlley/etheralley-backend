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

	metadata := &entities.NonFungibleMetadata{}
	err = json.Unmarshal([]byte(metadataString), &metadata)

	if err != nil {
		return nil, err
	}

	return metadata, err
}

func (g *Gateway) SaveNonFungibleMetadata(ctx context.Context, contract *entities.Contract, tokenId string, metadata *entities.NonFungibleMetadata) error {
	metadataBytes, err := json.Marshal(metadata)

	if err != nil {
		return err
	}

	_, err = g.client.Set(ctx, getFullKey(NFTNamespace, contract.Address, tokenId, contract.Blockchain), string(metadataBytes), time.Hour*24).Result()

	return err
}
