package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/core/entities"
)

func NewGetAllNonFungibleTokens(
	logger common.ILogger,
	getNonFungibleToken IGetNonFungibleTokenUseCase,
) IGetAllNonFungibleTokensUseCase {
	return &getAllNonFungibleTokensUseCase{
		logger,
		getNonFungibleToken,
	}
}

type getAllNonFungibleTokensUseCase struct {
	logger              common.ILogger
	getNonFungibleToken IGetNonFungibleTokenUseCase
}

type IGetAllNonFungibleTokensUseCase interface {
	// Get the metadata and balance of a slice of nfts.
	// Invalid contracts have a zeroed balance and nil metadata returned.
	Do(ctx context.Context, input *GetAllNonFungibleTokensInput) *[]entities.NonFungibleToken
}

type GetAllNonFungibleTokensInput struct {
	NonFungibleTokens *[]GetNonFungibleTokenInput `validate:"required"`
}

func (uc *getAllNonFungibleTokensUseCase) Do(ctx context.Context, input *GetAllNonFungibleTokensInput) *[]entities.NonFungibleToken {
	if err := common.ValidateStruct(input); err != nil {
		return &[]entities.NonFungibleToken{}
	}

	var wg sync.WaitGroup
	nfts := make([]entities.NonFungibleToken, len(*input.NonFungibleTokens))

	for i, n := range *input.NonFungibleTokens {
		wg.Add(1)

		go func(i int, nftInput GetNonFungibleTokenInput) {
			defer wg.Done()

			nft, err := uc.getNonFungibleToken.Do(ctx, &nftInput)

			if err != nil {
				nfts[i] = entities.NonFungibleToken{
					TokenId: nftInput.NonFungibleToken.TokenId,
					Contract: &entities.Contract{
						Blockchain: nftInput.NonFungibleToken.Contract.Blockchain,
						Address:    nftInput.NonFungibleToken.Contract.Address,
						Interface:  nftInput.NonFungibleToken.Contract.Interface,
					},
					Balance:  nil,
					Metadata: nil,
				}
			} else {
				nfts[i] = *nft
			}
		}(i, n)
	}

	wg.Wait()

	return &nfts
}
