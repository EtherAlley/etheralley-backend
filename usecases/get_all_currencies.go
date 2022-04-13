package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

type GetAllCurrenciesInput struct {
	Currencies *[]GetCurrencyInput `validate:"required"`
}

// Get the balances for the given list of addresses and blockchains
type IGetAllCurrenciesUseCase func(ctx context.Context, input *GetAllCurrenciesInput) *[]entities.Currency

func NewGetAllCurrenciesUseCase(
	logger common.ILogger,
	getCurrency IGetCurrencyUseCase,
) IGetAllCurrenciesUseCase {
	return func(ctx context.Context, input *GetAllCurrenciesInput) *[]entities.Currency {
		if err := common.ValidateStruct(input); err != nil {
			return &[]entities.Currency{}
		}

		var wg sync.WaitGroup
		currencies := make([]entities.Currency, len(*input.Currencies))

		for i, c := range *input.Currencies {
			wg.Add(1)

			go func(i int, currenycInput GetCurrencyInput) {
				defer wg.Done()

				currency, err := getCurrency(ctx, &currenycInput)

				if err != nil {
					currencies[i] = entities.Currency{
						Blockchain: currenycInput.Blockchain,
						Balance:    nil,
					}
				} else {
					currencies[i] = *currency
				}
			}(i, c)
		}

		wg.Wait()

		return &currencies
	}
}
