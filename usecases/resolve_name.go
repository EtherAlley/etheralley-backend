package usecases

import (
	"context"
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type ResolveENSNameInput struct {
	Address string `validate:"required,eth_addr"`
}

// resolve an ens name for an address
type IResolveENSNameUseCase func(ctx context.Context, input *ResolveENSNameInput) (name string, err error)

// Provided address is normalized to avoid user error
//
// Resolved values are cached
func NewResolveENSName(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	cacheGateway gateways.ICacheGateway,
) IResolveENSNameUseCase {
	return func(ctx context.Context, input *ResolveENSNameInput) (string, error) {
		if err := common.ValidateStruct(input); err != nil {
			return "", err
		}

		address := strings.ToLower(input.Address)

		name, err := cacheGateway.GetENSNameFromAddress(ctx, address)

		if err == nil {
			logger.Debugf(ctx, "cache hit for address %v -> ens name %v", address, name)
			return name, err
		}

		logger.Debugf(ctx, "cache miss getting ens name from address %v", address)

		name, err = blockchainGateway.GetENSNameFromAddress(ctx, address)

		cacheGateway.SaveENSName(ctx, address, name) // We should cache the result no matter what. Even the fact that they don't have an ens name

		return name, err
	}
}
