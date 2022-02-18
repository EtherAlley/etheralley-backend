package usecases

import (
	"context"
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type ResolveAddressInput struct {
	Value string `validate:"required"`
}

// resolve an address from an ens name
type IResolveAddressUseCase func(ctx context.Context, input *ResolveAddressInput) (address string, err error)

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
	return func(ctx context.Context, input *ResolveAddressInput) (string, error) {
		if err := common.ValidateStruct(input); err != nil {
			return "", err
		}

		value := strings.ToLower(input.Value)

		if strings.Contains(value, ".") {

			address, err := cacheGateway.GetENSAddressFromName(ctx, value)

			if err == nil {
				logger.Debugf(ctx, "cache hit for ens name %v -> address %v", input, address)
				return address, err
			}

			logger.Debugf(ctx, "cache miss for getting address from ens name %v", input)

			address, err = blockchainGateway.GetENSAddressFromName(ctx, value)

			if err != nil {
				logger.Debugf(ctx, "err getting address from ens name %v %v", input, err)
				return address, err
			}

			logger.Debugf(ctx, "chain hit for ens name %v -> address %v", input, address)

			address = strings.ToLower(address)

			cacheGateway.SaveENSAddress(ctx, value, address)

			return address, err
		}

		if err := common.ValidateField(value, `required,eth_addr`); err != nil {
			return value, err
		}

		return value, nil
	}
}
