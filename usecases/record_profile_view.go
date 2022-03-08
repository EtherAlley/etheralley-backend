package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type RecordProfileViewInput struct {
	Address   string ``
	IpAddress string ``
}

type IRecordProfileViewUseCase func(ctx context.Context, input *RecordProfileViewInput) error

func NewRecordProfileViewUseCase(logger common.ILogger, cacheGateway gateways.ICacheGateway) IRecordProfileViewUseCase {
	return func(ctx context.Context, input *RecordProfileViewInput) error {
		return cacheGateway.RecordAddressView(ctx, input.Address, input.IpAddress)
	}
}
