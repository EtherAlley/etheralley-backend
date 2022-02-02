package usecases

import (
	"context"
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

// if the address contains a ".", we assume its an attempted ens address and try to resolve
// attempt to cache resolved addresses
// if no "." we check for valid hex address and return if valid
// resolved addresses are lowered to accomodate user error
func NewResolveAddress(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	cacheGateway gateways.ICacheGateway,
) IResolveAddressUseCase {
	return func(ctx context.Context, input string) (string, error) {
		normalized := strings.ToLower(input)

		if strings.Contains(normalized, ".") {

			address, err := cacheGateway.GetENSAddressFromName(ctx, normalized)

			if err == nil {
				logger.Debugf("cache hit for ens name %v -> address %v", input, address)
				return address, err
			}

			logger.Debugf("cache miss for ens name %v", normalized)

			address, err = blockchainGateway.GetENSAddressFromName(normalized)

			if err != nil {
				logger.Debugf("chain miss for ens name %v err: %v", normalized, err)
				return address, err
			}

			address = strings.ToLower(address)

			logger.Debugf("chain hit for ens name %v -> address %v", normalized, address)

			cacheGateway.SaveENSAddress(ctx, normalized, address)

			return address, err
		}

		if err := common.ValidateField(normalized, `required,eth_addr`); err != nil {
			return normalized, err
		}

		return normalized, nil
	}
}
