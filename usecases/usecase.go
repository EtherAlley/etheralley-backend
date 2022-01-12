package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/entities"
)

// get the profile for the provided address
type GetProfileUseCase func(ctx context.Context, address string) (*entities.Profile, error)

// save the provided profile
type SaveProfileUseCase func(ctx context.Context, profile *entities.Profile) error

// get a challenge message for the provided address
type GetChallengeUseCase func(ctx context.Context, address string) (*entities.Challenge, error)

// verify if the provided signature was signed with the correct address and signed the correct challenge message
type VerifyChallengeUseCase func(ctx context.Context, address string, sigHex string) error

// get the metadata and balance of an nft
type GetNonFungibleTokenUseCase func(ctx context.Context, address string, contract *entities.Contract, tokenId string) (*entities.NonFungibleToken, error)

//get the metadata and balance of a slice of nfts
type GetAllNonFungibleTokensUseCase func(ctx context.Context, address string, partials *[]entities.NonFungibleToken) *[]entities.NonFungibleToken

// validate an address and resolve from ens name if provided
type GetValidAddressUseCase func(ctx context.Context, address string) (string, error)

// get the metadata and balance of an nft
type GetFungibleTokenUseCase func(ctx context.Context, address string, contract *entities.Contract) (*entities.FungibleToken, error)

// get a slice of fungible tokens for the given address and contracts
type GetAllFungibleTokensUseCase func(ctx context.Context, address string, contract *[]entities.Contract) *[]entities.FungibleToken
