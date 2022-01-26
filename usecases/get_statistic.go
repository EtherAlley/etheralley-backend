package usecases

import (
	"context"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/thegraph"
)

func NewGetStatisticUseCase(logger *common.Logger, blockchainIndexGateway *thegraph.Gateway) GetStatisticUseCase {
	return GetStatistic(logger, blockchainIndexGateway)
}

func GetStatistic(logger *common.Logger, blockchainIndexGateway gateways.IBlockchainIndexGateway) GetStatisticUseCase {
	return func(ctx context.Context, address string, contract *entities.Contract) (*entities.Statistic, error) {
		if err := common.ValidateStruct(contract); err != nil {
			return nil, err
		}

		// TODO: call gateway function based on interfaces
		data, err := blockchainIndexGateway.GetSwaps(ctx, address, contract)

		if err != nil {
			return nil, err
		}

		return &entities.Statistic{
			Contract: contract,
			Data:     data,
		}, nil
	}
}
