package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/entities"
)

// get the profile for the provided address
type GetProfileUsecase func(ctx context.Context, address string) (*entities.Profile, error)

// save the provided profile
type SaveProfileUseCase func(ctx context.Context, profile *entities.Profile) error

// get a challenge message for the provided address
type GetChallengeUseCase func(ctx context.Context, address string) (*entities.Challenge, error)

// verify if the provided signature was signed with the correct address and signed the correct challenge message
type VerifyChallengeUseCase func(ctx context.Context, address string, sigHex string) error

// get the metadata and ownership of an nft
type GetNFTUseCase func(ctx context.Context, address string, blockchain string, contractAddress string, schema_name string, token_id string) (*entities.NFT, error)

//
type HydrateNFTsUseCase func(ctx context.Context, address string, partialNFTs []entities.NFT) []entities.NFT
