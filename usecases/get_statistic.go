package usecases

import (
	"context"
	"errors"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
	"github.com/etheralley/etheralley-core-api/gateways"
)

type GetStatisticsInput struct {
	Address   string          `validate:"required,eth_addr"`
	Statistic *StatisticInput `validate:"required,dive"`
}

// get the statistic for a given address and contract
type IGetStatisticUseCase func(ctx context.Context, input *GetStatisticsInput) (*entities.Statistic, error)

func NewGetStatistic(
	logger common.ILogger,
	blockchainIndexGateway gateways.IBlockchainIndexGateway,
) IGetStatisticUseCase {
	return func(ctx context.Context, input *GetStatisticsInput) (*entities.Statistic, error) {
		if err := common.ValidateStruct(input); err != nil {
			return nil, err
		}

		address := input.Address
		statType := input.Statistic.Type
		contract := &entities.Contract{
			Blockchain: input.Statistic.Contract.Blockchain,
			Address:    input.Statistic.Contract.Address,
			Interface:  input.Statistic.Contract.Interface,
		}

		var data interface{}
		var err error
		switch statType {
		case common.SWAP:
			data, err = blockchainIndexGateway.GetSwaps(ctx, address, contract)
		case common.STAKE:
			data, err = blockchainIndexGateway.GetStake(ctx, address, contract)
		default:
			data, err = nil, errors.New("invalid stat type")
		}

		if err != nil {
			logger.Info(ctx).Err(err).Msgf("err getting statistic %v %v %v %v", address, statType, contract.Blockchain, contract.Interface)
			return nil, err
		}

		return &entities.Statistic{
			Type:     statType,
			Contract: contract,
			Data:     data,
		}, nil
	}
}
