package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

func NewGetAllFungibleTokens(
	logger common.ILogger,
	getFungibleToken IGetFungibleTokenUseCase,
) IGetAllFungibleTokensUseCase {
	return &getAllFungibleTokens{
		logger,
		getFungibleToken,
	}
}

type getAllFungibleTokens struct {
	logger           common.ILogger
	getFungibleToken IGetFungibleTokenUseCase
}

type GetAllFungibleTokensInput struct {
	Tokens *[]GetFungibleTokenInput `validate:"required"`
}

type IGetAllFungibleTokensUseCase interface {
	// Get a slice of fungible tokens for the given contracts/address.
	// Invalid contracts will return a token with a zeroed balance.
	Do(ctx context.Context, input *GetAllFungibleTokensInput) *[]entities.FungibleToken
}

func (uc *getAllFungibleTokens) Do(ctx context.Context, input *GetAllFungibleTokensInput) *[]entities.FungibleToken {
	if err := common.ValidateStruct(input); err != nil {
		return &[]entities.FungibleToken{}
	}

	var wg sync.WaitGroup

	tokens := make([]entities.FungibleToken, len(*input.Tokens))

	for i, t := range *input.Tokens {
		wg.Add(1)

		go func(i int, tokenInput GetFungibleTokenInput) {
			defer wg.Done()

			token, err := uc.getFungibleToken.Do(ctx, &tokenInput)

			if err != nil {
				tokens[i] = entities.FungibleToken{
					Contract: &entities.Contract{
						Blockchain: tokenInput.Token.Contract.Blockchain,
						Address:    tokenInput.Token.Contract.Address,
						Interface:  tokenInput.Token.Contract.Interface,
					},
					Balance: nil,
					Metadata: &entities.FungibleMetadata{
						Name:     nil,
						Symbol:   nil,
						Decimals: nil,
					},
				}
			} else {
				tokens[i] = *token
			}

		}(i, t)
	}

	wg.Wait()

	return &tokens
}
