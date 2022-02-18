package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

type GetAllNonFungibleTokensInput struct {
	NonFungibleTokens *[]GetNonFungibleTokenInput `validate:"required"`
}

//Get the metadata and balance of a slice of nfts
//
// For transient info for each nft provided concurrently
//
// Invalid contracts are discarded
type IGetAllNonFungibleTokensUseCase func(ctx context.Context, input *GetAllNonFungibleTokensInput) *[]entities.NonFungibleToken

func NewGetAllNonFungibleTokens(
	logger common.ILogger,
	getNonFungibleToken IGetNonFungibleTokenUseCase,
) IGetAllNonFungibleTokensUseCase {
	return func(ctx context.Context, input *GetAllNonFungibleTokensInput) *[]entities.NonFungibleToken {
		if err := common.ValidateStruct(input); err != nil {
			return &[]entities.NonFungibleToken{}
		}

		var wg sync.WaitGroup

		nfts := make([]*entities.NonFungibleToken, len(*input.NonFungibleTokens))
		for i, n := range *input.NonFungibleTokens {
			wg.Add(1)

			go func(i int, nftInput GetNonFungibleTokenInput) {
				defer wg.Done()

				nft, err := getNonFungibleToken(ctx, &nftInput)

				if err != nil {
					logger.Errf(ctx, err, "invalid nft: contract address %v token id %v chain %v", nftInput.NonFungibleToken.Contract.Address, nftInput.NonFungibleToken.TokenId, nftInput.NonFungibleToken.Contract.Blockchain)
					return
				}

				nfts[i] = nft

			}(i, n)
		}

		wg.Wait()

		trimmedNfts := []entities.NonFungibleToken{}
		for _, nft := range nfts {
			if nft != nil {
				trimmedNfts = append(trimmedNfts, *nft)
			}
		}

		return &trimmedNfts
	}
}
