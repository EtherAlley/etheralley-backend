package gateways

import (
	"context"
	"math/big"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/core/entities"
)

type IDatabaseGateway interface {
	Init(context.Context) error

	GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error)
	SaveProfile(ctx context.Context, profile *entities.Profile) error
}

type ICacheGateway interface {
	Init(context.Context) error

	GetProfileByAddress(ctx context.Context, address string) (*entities.Profile, error)
	SaveProfile(ctx context.Context, profile *entities.Profile) error
	DeleteProfile(ctx context.Context, address string) error

	GetChallengeByAddress(ctx context.Context, address string) (*entities.Challenge, error)
	SaveChallenge(ctx context.Context, challenge *entities.Challenge) error

	GetNonFungibleMetadata(ctx context.Context, contract *entities.Contract, tokenId string) (*entities.NonFungibleMetadata, error)
	SaveNonFungibleMetadata(ctx context.Context, contract *entities.Contract, tokenId string, metadata *entities.NonFungibleMetadata) error

	GetFungibleMetadata(ctx context.Context, contract *entities.Contract) (*entities.FungibleMetadata, error)
	SaveFungibleMetadata(ctx context.Context, contract *entities.Contract, metadata *entities.FungibleMetadata) error

	GetENSAddressFromName(ctx context.Context, name string) (string, error)
	SaveENSAddress(ctx context.Context, name string, address string) error

	GetENSNameFromAddress(ctx context.Context, address string) (string, error)
	SaveENSName(ctx context.Context, address string, name string) error

	RecordAddressView(ctx context.Context, address string, ipAddress string) error
	GetTopViewedAddresses(ctx context.Context) (*[]string, error)
	GetTopViewedProfiles(ctx context.Context) (*[]entities.Profile, error)
	SaveTopViewedProfiles(ctx context.Context, profiles *[]entities.Profile) error

	GetStoreListings(ctx context.Context, tokenIds *[]string) (*[]entities.Listing, error)
	SaveStoreListings(ctx context.Context, listings *[]entities.Listing) error

	VerifyRateLimit(ctx context.Context, ipAddress string) error
}

type IBlockchainGateway interface {
	GetAccountBalance(ctx context.Context, blockchain common.Blockchain, address string) (string, error)

	GetERC1155URI(ctx context.Context, contract *entities.Contract, tokenId string) (string, error)
	GetERC721URI(ctx context.Context, contract *entities.Contract, tokenId string) (string, error)
	GetERC1155Balance(ctx context.Context, address string, contract *entities.Contract, tokenId string) (string, error)
	GetERC721Balance(ctx context.Context, address string, contract *entities.Contract, tokenId string) (string, error)
	GetPunkBalance(ctx context.Context, address string, contract *entities.Contract, tokenId string) (string, error)

	GetERC20Balance(ctx context.Context, address string, contract *entities.Contract) (string, error)
	GetERC20Name(ctx context.Context, contract *entities.Contract) (string, error)
	GetERC20Symbol(ctx context.Context, contract *entities.Contract) (string, error)
	GetERC20Decimals(ctx context.Context, contract *entities.Contract) (uint8, error)

	GetENSAddressFromName(ctx context.Context, ensName string) (string, error)
	GetENSNameFromAddress(ctx context.Context, address string) (name string, err error)

	GetTransactionData(ctx context.Context, tx *entities.Transaction) (*entities.TransactionData, error)

	GetStoreListingInfo(ctx context.Context, ids *[]string) (*[]entities.ListingInfo, error)
	GetStoreBalanceBatch(ctx context.Context, address string, ids *[]string) ([]*big.Int, error)
}

type IBlockchainIndexGateway interface {
	GetSwaps(ctx context.Context, address string, contract *entities.Contract) (interface{}, error)
	GetStake(ctx context.Context, address string, contract *entities.Contract) (interface{}, error)
}

type IOffchainGateway interface {
	Init(context.Context) error

	GetNonFungibleTokens(ctx context.Context, address string) (*[]entities.NonFungibleToken, error)
	GetNonFungibleMetadata(ctx context.Context, uri string) (*entities.NonFungibleMetadata, error)
	GetPunkMetadata(ctx context.Context, tokenId string) (*entities.NonFungibleMetadata, error)
	GetKittieMetadata(ctx context.Context, tokenId string) (*entities.NonFungibleMetadata, error)

	GetFungibleMetadata(ctx context.Context, contract *entities.Contract) (*entities.FungibleMetadata, error)
}
