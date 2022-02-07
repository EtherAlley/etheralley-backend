package usecases

import (
	"context"
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

// Provided address is normalized to avoid user error
//
// Resolved values are cached
func NewResolveENSName(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	cacheGateway gateways.ICacheGateway,
) IResolveENSNameUseCase {
	return func(ctx context.Context, address string) (string, error) {
		address = strings.ToLower(address)

		name, err := cacheGateway.GetENSNameFromAddress(ctx, address)

		if err == nil {
			logger.Debugf(ctx, "cache hit for address %v -> ens name %v", address, name)
			return name, err
		}

		logger.Debugf(ctx, "cache miss getting ens name from address %v", address)

		return blockchainGateway.GetENSNameFromAddress(ctx, address)
	}
}
