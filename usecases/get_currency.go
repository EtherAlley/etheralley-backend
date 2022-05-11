package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type GetCurrencyInput struct {
	Address    string            `validate:"required,eth_addr"`
	Blockchain common.Blockchain `validate:"required,oneof=ethereum polygon arbitrum optimism"`
}

// Get the balance for a given address and blockchain
type IGetCurrencyUseCase func(ctx context.Context, input *GetCurrencyInput) (*entities.Currency, error)

func NewGetCurrency(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
) IGetCurrencyUseCase {
	return func(ctx context.Context, input *GetCurrencyInput) (*entities.Currency, error) {
		if err := common.ValidateStruct(input); err != nil {
			return nil, err
		}

		balance, err := blockchainGateway.GetAccountBalance(ctx, input.Blockchain, input.Address)

		if err != nil {
			logger.Info(ctx).Err(err).Msgf("err getting account balance %v %v", input.Blockchain, input.Address)
			return nil, err
		}

		return &entities.Currency{
			Blockchain: input.Blockchain,
			Balance:    &balance,
		}, nil
	}
}
