package usecases

import (
	"context"
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

// Attempts to detect provided input format and resolve ens address
//
// Provided input is normalized to avoid user error
//
// Output address is also normalized for consistency
//
// If the address contains a ".", we assume its an attempted ens address and try to resolve.
// Otherwise, if input is in valid hex format we gracefully return it without err.
//
// Resolved values are cached
func NewResolveENSAddress(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	cacheGateway gateways.ICacheGateway,
) IResolveAddressUseCase {
	return func(ctx context.Context, input string) (string, error) {
		input = strings.ToLower(input)

		if strings.Contains(input, ".") {

			address, err := cacheGateway.GetENSAddressFromName(ctx, input)

			if err == nil {
				logger.Debugf(ctx, "cache hit for ens name %v -> address %v", input, address)
				return address, err
			}

			logger.Debugf(ctx, "cache miss for getting address from ens name %v", input)

			address, err = blockchainGateway.GetENSAddressFromName(ctx, input)

			if err != nil {
				logger.Debugf(ctx, "err getting address from ens name %v %v", input, err)
				return address, err
			}

			logger.Debugf(ctx, "chain hit for ens name %v -> address %v", input, address)

			address = strings.ToLower(address)

			cacheGateway.SaveENSAddress(ctx, input, address)

			return address, err
		}

		if err := common.ValidateField(input, `required,eth_addr`); err != nil {
			return input, err
		}

		return input, nil
	}
}
