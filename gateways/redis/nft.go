package redis

import (
	"context"
	"encoding/json"

	"github.com/etheralley/etheralley-core-api/entities"
)

const NFTNamespace = "nft"

func (g *Gateway) GetNFTMetadata(ctx context.Context, location *entities.NFTLocation) (*entities.NFTMetadata, error) {

	metadataString, err := g.client.Get(ctx, GetFullKey(NFTNamespace, location.ContractAddress, location.TokenId)).Result()

	if err != nil {
		return nil, err
	}

	metadata := &entities.NFTMetadata{}
	err = json.Unmarshal([]byte(metadataString), &metadata)

	if err != nil {
		return nil, err
	}

	return metadata, err
}

func (g *Gateway) SaveNFTMetadata(ctx context.Context, location *entities.NFTLocation, metadata *entities.NFTMetadata) error {
	metadataBytes, err := json.Marshal(metadata)

	if err != nil {
		return err
	}

	_, err = g.client.Set(ctx, GetFullKey(NFTNamespace, location.ContractAddress, location.TokenId), string(metadataBytes), 0).Result()

	return err
}
