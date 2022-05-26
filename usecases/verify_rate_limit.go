package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

func NewVerifyRateLimit(
	logger common.ILogger,
	cacheGateway gateways.ICacheGateway,
) IVerifyRateLimitUseCase {
	return &verifyRateLimitUseCase{
		logger,
		cacheGateway,
	}
}

type verifyRateLimitUseCase struct {
	logger       common.ILogger
	cacheGateway gateways.ICacheGateway
}

type IVerifyRateLimitUseCase interface {
	// Verify that the provided IpAddress is not rate limited
	Do(ctx context.Context, input *VerifyRateLimitInput) error
}

type VerifyRateLimitInput struct {
	IpAddress string `validate:"required,ip"`
}

func (uc *verifyRateLimitUseCase) Do(ctx context.Context, input *VerifyRateLimitInput) error {
	if err := common.ValidateStruct(input); err != nil {
		return err
	}

	if err := uc.cacheGateway.VerifyRateLimit(ctx, input.IpAddress); err != nil {
		return err
	}

	return nil
}
