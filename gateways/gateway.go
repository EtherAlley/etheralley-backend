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
	GetNFTMetadata(contractAddress string, tokenId string, schemaName string) (*entities.NFTMetadata, error)
	VerifyOwner(contractAddress string, address string, tokenId string, schemaName string) (bool, error)
}

type INFTAPIGateway interface {
	GetNFTs(address string) ([]entities.NFT, error)
}
