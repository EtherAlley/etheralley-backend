package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/core/entities"
	"github.com/etheralley/etheralley-core-api/core/gateways"
)

func NewGetCurrency(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
) IGetCurrencyUseCase {
	return &getCurrencyUseCase{
		logger,
		blockchainGateway,
	}
}

type getCurrencyUseCase struct {
	logger            common.ILogger
	blockchainGateway gateways.IBlockchainGateway
}

type IGetCurrencyUseCase interface {
	// Get the balance for a given address and blockchain
	Do(ctx context.Context, input *GetCurrencyInput) (*entities.Currency, error)
}

type GetCurrencyInput struct {
	Address    string            `validate:"required,eth_addr"`
	Blockchain common.Blockchain `validate:"required,oneof=ethereum polygon arbitrum optimism"`
}

func (uc *getCurrencyUseCase) Do(ctx context.Context, input *GetCurrencyInput) (*entities.Currency, error) {
	if err := common.ValidateStruct(input); err != nil {
		return nil, err
	}

	balance, err := uc.blockchainGateway.GetAccountBalance(ctx, input.Blockchain, input.Address)

	if err != nil {
		uc.logger.Info(ctx).Err(err).Msgf("err getting account balance %v %v", input.Blockchain, input.Address)
		return nil, err
	}

	return &entities.Currency{
		Blockchain: input.Blockchain,
		Balance:    &balance,
	}, nil
}
