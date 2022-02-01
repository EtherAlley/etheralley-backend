package usecases

import (
	"context"
	"errors"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
	"github.com/etheralley/etheralley-core-api/gateways/thegraph"
)

func NewGetStatisticUseCase(logger *common.Logger, blockchainIndexGateway *thegraph.Gateway) IGetStatisticUseCase {
	return GetStatistic(logger, blockchainIndexGateway)
}

func GetStatistic(logger *common.Logger, blockchainIndexGateway gateways.IBlockchainIndexGateway) IGetStatisticUseCase {
	return func(ctx context.Context, address string, contract *entities.Contract, statType common.StatisticType) (*entities.Statistic, error) {
		if err := common.ValidateStruct(contract); err != nil {
			return nil, err
		}

		var data interface{}
		var err error
		switch statType {
		case common.SWAP:
			data, err = blockchainIndexGateway.GetSwaps(ctx, address, contract)
		default:
			data, err = nil, errors.New("invalid stat type")
		}

		if err != nil {
			return nil, err
		}

		return &entities.Statistic{
			Type:     statType,
			Contract: contract,
			Data:     data,
		}, nil
	}
}
