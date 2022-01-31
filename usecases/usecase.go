package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/entities"
)

// get the profile for the provided address
type IGetProfileUseCase func(ctx context.Context, address string) (*entities.Profile, error)

// get a default profile for the provided address
type IGetDefaultProfileUseCase func(ctx context.Context, address string) (*entities.Profile, error)

// save the provided profile
type ISaveProfileUseCase func(ctx context.Context, profile *entities.Profile) error

// get a challenge message for the provided address
type IGetChallengeUseCase func(ctx context.Context, address string) (*entities.Challenge, error)

// verify if the provided signature was signed with the correct address and signed the correct challenge message
type IVerifyChallengeUseCase func(ctx context.Context, address string, sigHex string) error

// resolve an address from an ens name
type IResolveAddressUseCase func(ctx context.Context, ensName string) (string, error)

// get the metadata and balance of an nft
type IGetNonFungibleTokenUseCase func(ctx context.Context, address string, contract *entities.Contract, tokenId string) (*entities.NonFungibleToken, error)

//get the metadata and balance of a slice of nfts
type IGetAllNonFungibleTokensUseCase func(ctx context.Context, address string, partials *[]entities.NonFungibleToken) *[]entities.NonFungibleToken

// get the metadata and balance of an nft
type IGetFungibleTokenUseCase func(ctx context.Context, address string, contract *entities.Contract) (*entities.FungibleToken, error)

// get a slice of fungible tokens for the given address and contracts
type IGetAllFungibleTokensUseCase func(ctx context.Context, address string, contract *[]entities.Contract) *[]entities.FungibleToken

// get the statistic for a given address and contract
type IGetStatisticUseCase func(ctx context.Context, address string, contract *entities.Contract) (*entities.Statistic, error)

// get all statistics for a given address and slice of contract
type IGetAllStatisticsUseCase func(ctx context.Context, address string, contract *[]entities.Contract) *[]entities.Statistic
