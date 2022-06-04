package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/core/entities"
)

func NewGetAllCurrenciesUseCase(
	logger common.ILogger,
	getCurrency IGetCurrencyUseCase,
) IGetAllCurrenciesUseCase {
	return &getAllCurrenciesUsecase{
		logger,
		getCurrency,
	}
}

type getAllCurrenciesUsecase struct {
	logger      common.ILogger
	getCurrency IGetCurrencyUseCase
}

type IGetAllCurrenciesUseCase interface {
	// Get the balances for the given list of addresses and blockchains.
	// A nil balance will be returned for any invalid contracts
	Do(ctx context.Context, input *GetAllCurrenciesInput) *[]entities.Currency
}

type GetAllCurrenciesInput struct {
	Currencies *[]GetCurrencyInput `validate:"required"`
}

func (uc *getAllCurrenciesUsecase) Do(ctx context.Context, input *GetAllCurrenciesInput) *[]entities.Currency {
	if err := common.ValidateStruct(input); err != nil {
		return &[]entities.Currency{}
	}

	var wg sync.WaitGroup
	currencies := make([]entities.Currency, len(*input.Currencies))

	for i, c := range *input.Currencies {
		wg.Add(1)

		go func(i int, currenycInput GetCurrencyInput) {
			defer wg.Done()

			currency, err := uc.getCurrency.Do(ctx, &currenycInput)

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
