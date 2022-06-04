package usecases

import (
	"context"
	"strings"

	"github.com/etheralley/etheralley-apis/common"
	"github.com/etheralley/etheralley-apis/core/gateways"
)

func NewResolveENSName(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	cacheGateway gateways.ICacheGateway,
) IResolveENSNameUseCase {
	return &resolveENSNameUseCase{
		logger,
		blockchainGateway,
		cacheGateway,
	}
}

type resolveENSNameUseCase struct {
	logger            common.ILogger
	blockchainGateway gateways.IBlockchainGateway
	cacheGateway      gateways.ICacheGateway
}

type IResolveENSNameUseCase interface {
	// Resolve an ens name for an address
	Do(ctx context.Context, input *ResolveENSNameInput) (name string, err error)
}

type ResolveENSNameInput struct {
	Address string `validate:"required,eth_addr"`
}

// Provided address is normalized to avoid user error.
// Resolved values are cached.
func (uc *resolveENSNameUseCase) Do(ctx context.Context, input *ResolveENSNameInput) (string, error) {
	if err := common.ValidateStruct(input); err != nil {
		return "", err
	}

	address := strings.ToLower(input.Address)

	name, err := uc.cacheGateway.GetENSNameFromAddress(ctx, address)

	if err == nil {
		uc.logger.Debug(ctx).Msgf("cache hit for address %v -> ens name %v", address, name)
		return name, err
	}

	uc.logger.Debug(ctx).Msgf("cache miss getting ens name from address %v", address)

	name, err = uc.blockchainGateway.GetENSNameFromAddress(ctx, address)

	if err != nil {
		uc.logger.Info(ctx).Err(err).Msgf("err getting ens name from address %v", address)
	}

	uc.logger.Debug(ctx).Msgf("chain hit for address %v -> ens name %v", address, name)

	uc.cacheGateway.SaveENSName(ctx, address, name) // We should cache the result no matter what. Even the fact that they don't have an ens name

	return name, err
}
