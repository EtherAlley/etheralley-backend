package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type RecordProfileViewInput struct {
	Address   string `validate:"required,eth_addr"`
	IpAddress string `validate:"required,ip"`
}

type IRecordProfileViewUseCase func(ctx context.Context, input *RecordProfileViewInput) error

func NewRecordProfileViewUseCase(logger common.ILogger, cacheGateway gateways.ICacheGateway) IRecordProfileViewUseCase {
	return func(ctx context.Context, input *RecordProfileViewInput) error {
		if err := common.ValidateStruct(input); err != nil {
			return err
		}

		return cacheGateway.RecordAddressView(ctx, input.Address, input.IpAddress)
	}
}
