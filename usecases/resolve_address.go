package usecases

import (
	"context"
	"strings"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
)

func NewResolveENSAddress(
	logger common.ILogger,
	blockchainGateway gateways.IBlockchainGateway,
	cacheGateway gateways.ICacheGateway,
) IResolveAddressUseCase {
	return &resolveAddressUseCase{
		logger,
		blockchainGateway,
		cacheGateway,
	}
}

type resolveAddressUseCase struct {
	logger            common.ILogger
	blockchainGateway gateways.IBlockchainGateway
	cacheGateway      gateways.ICacheGateway
}

type IResolveAddressUseCase interface {
	// Resolve an address from an ens name
	Do(ctx context.Context, input *ResolveAddressInput) (string, error)
}

type ResolveAddressInput struct {
	Value string `validate:"required"`
}

// Attempts to detect provided input format and resolve ens address.
// Provided input is normalized to avoid user error.
// Output address is also normalized for consistency.
// If the address contains a ".", we assume its an attempted ens address and try to resolve.
// Otherwise, if input is in valid hex format we gracefully return it without err.
// Resolved values are cached
func (uc *resolveAddressUseCase) Do(ctx context.Context, input *ResolveAddressInput) (string, error) {
	if err := common.ValidateStruct(input); err != nil {
		return "", err
	}

	value := strings.ToLower(input.Value)

	if strings.Contains(value, ".") {

		address, err := uc.cacheGateway.GetENSAddressFromName(ctx, value)

		if err == nil {
			uc.logger.Debug(ctx).Msgf("cache hit for ens name %v -> address %v", input, address)
			return address, err
		}

		uc.logger.Debug(ctx).Msgf("cache miss for getting address from ens name %v", input)

		address, err = uc.blockchainGateway.GetENSAddressFromName(ctx, value)

		if err != nil {
			uc.logger.Info(ctx).Err(err).Msgf("err getting address from ens name %v", input)
			return address, err
		}

		uc.logger.Debug(ctx).Msgf("chain hit for ens name %v -> address %v", input, address)

		address = strings.ToLower(address)

		uc.cacheGateway.SaveENSAddress(ctx, value, address)

		return address, err
	}

	if err := common.ValidateField(value, `required,eth_addr`); err != nil {
		return value, err
	}

	return value, nil
}
