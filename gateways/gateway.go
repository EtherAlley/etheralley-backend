package gateways

import (
	"context"

	"github.com/etheralley/etheralley-core-api/entities"
)

type IDatabaseGateway interface {
	GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error)
	SaveProfile(ctx context.Context, profile *entities.Profile) error
}

type ICacheGateway interface {
	GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error)
	SaveProfile(ctx context.Context, profile *entities.Profile) error
	GetChallengeByAddress(ctx context.Context, address string) (*entities.Challenge, error)
	SaveChallenge(ctx context.Context, challenge *entities.Challenge) error
}

type IBlockchainGateway interface {
	GetERC1155NFTMetadata(contractAddress string, tokenId string) (*NFTMetadata, error)
	GetERC721NFTMetadata(contractAddress string, tokenId string) (*NFTMetadata, error)
	VerifyERC1155Owner(contractAddress string, address string, tokenId string) (bool, error)
	VerifyERC721Owner(contractAddress string, address string, tokenId string) (bool, error)
}

// https://eips.ethereum.org/EIPS/eip-1155
// https://eips.ethereum.org/EIPS/eip-721
type NFTMetadata struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Image       string                 `json:"image"`
	Attributes  []map[string]string    `bson:"attributes" json:"attributes"`
	Properties  map[string]interface{} `bson:"properties" json:"properties"`
}
