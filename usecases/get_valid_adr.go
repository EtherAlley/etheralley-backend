package usecases

import (
	"context"
	"errors"
	"strings"

	eaCommon "github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/ethereum"
	"github.com/etheralley/etheralley-core-api/gateways/redis"
	"github.com/ethereum/go-ethereum/common"
)

func NewGetValidAddressUseCase(logger *eaCommon.Logger, blockchainGateway *ethereum.Gateway, cacheGateway *redis.Gateway) GetValidAddressUseCase {
	return GetValidAddress(logger, blockchainGateway, cacheGateway)
}

// TODO: casing???
// if the address contains a ".", we assume its an attempted ens address and try to resolve
// we first check the cache
// if no "." we check for valid hex address and return if so
func GetValidAddress(logger *eaCommon.Logger, blockchainGateway gateways.IBlockchainGateway, cacheGateway gateways.ICacheGateway) GetValidAddressUseCase {
	return func(ctx context.Context, input string) (address string, err error) {

		if input == "" {
			return address, errors.New("empty input")
		}

		if strings.Contains(input, ".") {
			address, err = cacheGateway.GetENSAddressFromName(ctx, input)

			if err == nil {
				logger.Debugf("cache hit for ens name %v -> address %v", input, address)
				return
			}

			logger.Debugf("cache miss for ens name %v", input)

			address, err = blockchainGateway.GetENSAddressFromName(input)

			if err != nil {
				logger.Debugf("chain miss for ens name %v err: %v", input, err)
				return
			}

			logger.Debugf("chain hit for ens name %v -> address %v", input, address)

			cacheGateway.SaveENSAddress(ctx, input, address)

			return
		}

		if common.IsHexAddress(input) {
			return input, nil
		}

		return address, errors.New("invalida hex format")
	}
}
