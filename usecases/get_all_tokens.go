package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

type GetAllFungibleTokensInput struct {
	Tokens *[]GetFungibleTokenInput `validate:"required"`
}

// Get a slice of fungible tokens for the given contracts/address
//
// Fetch the full token info for each contract provided
//
// Invalid contracts are discarded
type IGetAllFungibleTokensUseCase func(ctx context.Context, input *GetAllFungibleTokensInput) *[]entities.FungibleToken

func NewGetAllFungibleTokens(
	logger common.ILogger,
	getFungibleToken IGetFungibleTokenUseCase,
) IGetAllFungibleTokensUseCase {
	return func(ctx context.Context, input *GetAllFungibleTokensInput) *[]entities.FungibleToken {
		if err := common.ValidateStruct(input); err != nil {
			return &[]entities.FungibleToken{}
		}

		var wg sync.WaitGroup

		arr := make([]*entities.FungibleToken, len(*input.Tokens))

		for i, t := range *input.Tokens {
			wg.Add(1)

			go func(i int, tokenInput GetFungibleTokenInput) {
				defer wg.Done()

				token, err := getFungibleToken(ctx, &tokenInput)

				if err != nil {
					logger.Errf(ctx, err, "invalid token: contract address %v chain %v interface %v", tokenInput.Token.Contract.Address, tokenInput.Token.Contract.Blockchain, tokenInput.Token.Contract.Interface)
					return
				}

				arr[i] = token
			}(i, t)
		}

		wg.Wait()

		trimmedTokens := []entities.FungibleToken{}
		for _, token := range arr {
			if token != nil {
				trimmedTokens = append(trimmedTokens, *token)
			}
		}

		return &trimmedTokens
	}
}
