package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

func NewGetAllFungibleTokensUseCase(logger *common.Logger, getFungibleToken IGetFungibleTokenUseCase) IGetAllFungibleTokensUseCase {
	return GetAllFungibleTokens(logger, getFungibleToken)
}

// fetch the full token info for each contract provided
// we can use a simple slice here since each result in the go routine writes to its own index location
// invalid contracts are discarded
func GetAllFungibleTokens(logger *common.Logger, getFungibleToken IGetFungibleTokenUseCase) IGetAllFungibleTokensUseCase {
	return func(ctx context.Context, address string, contracts *[]entities.Contract) *[]entities.FungibleToken {
		var wg sync.WaitGroup

		arr := make([]*entities.FungibleToken, len(*contracts))

		for i, contract := range *contracts {
			wg.Add(1)

			go func(i int, c entities.Contract) {
				defer wg.Done()

				token, err := getFungibleToken(ctx, address, &c)

				if err != nil {
					logger.Errf(err, "invalid token: contract address %v chain %v", c.Address, c.Blockchain)
					return
				}

				arr[i] = token
			}(i, contract)
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
