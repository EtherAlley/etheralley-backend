package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

func NewRecordProfileViewUseCase(logger common.ILogger, cacheGateway gateways.ICacheGateway) IRecordProfileViewUseCase {
	return &recordProfileViewUseCase{
		logger,
		cacheGateway,
	}
}

type recordProfileViewUseCase struct {
	logger       common.ILogger
	cacheGateway gateways.ICacheGateway
}

type IRecordProfileViewUseCase interface {
	Do(ctx context.Context, input *RecordProfileViewInput) error
}

type RecordProfileViewInput struct {
	Address   string `validate:"required,eth_addr"`
	IpAddress string `validate:"required,ip"`
}

func (uc *recordProfileViewUseCase) Do(ctx context.Context, input *RecordProfileViewInput) error {
	if err := common.ValidateStruct(input); err != nil {
		uc.logger.Info(ctx).Err(err).Msgf("err recording profile view with ip %v", input.IpAddress)
		return err
	}

	return uc.cacheGateway.RecordAddressView(ctx, input.Address, input.IpAddress)
}
