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
	GetNFTMetadata(ctx context.Context, location *entities.NFTLocation) (*entities.NFTMetadata, error)
	SaveNFTMetadata(ctx context.Context, location *entities.NFTLocation, metadata *entities.NFTMetadata) error
}

type IBlockchainGateway interface {
	GetNFTMetadata(location *entities.NFTLocation) (*entities.NFTMetadata, error)
	VerifyOwner(address string, location *entities.NFTLocation) (bool, error)
}

type INFTAPIGateway interface {
	GetNFTs(address string) (*[]entities.NFT, error)
}
