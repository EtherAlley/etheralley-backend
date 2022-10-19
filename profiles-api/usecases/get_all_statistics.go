package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-backend/common"
	"github.com/etheralley/etheralley-backend/profiles-api/entities"
)

func NewGetAllStatistics(
	logger common.ILogger,
	getStatistic IGetStatisticUseCase,
) IGetAllStatisticsUseCase {
	return &getAllStatisticsUseCase{
		logger,
		getStatistic,
	}
}

type getAllStatisticsUseCase struct {
	logger       common.ILogger
	getStatistic IGetStatisticUseCase
}

type IGetAllStatisticsUseCase interface {
	// Get all statistic data for a given slice of statistics.
	// Invalid contracts will return a statistic with nil Data.
	Do(ctx context.Context, input *GetAllStatisticsInput) *[]entities.Statistic
}

type GetAllStatisticsInput struct {
	Stats *[]GetStatisticsInput `validate:"required"`
}

func (uc *getAllStatisticsUseCase) Do(ctx context.Context, input *GetAllStatisticsInput) *[]entities.Statistic {
	if err := common.ValidateStruct(input); err != nil {
		return &[]entities.Statistic{}
	}

	var wg sync.WaitGroup

	stats := make([]entities.Statistic, len(*input.Stats))

	for i, s := range *input.Stats {
		wg.Add(1)

		go func(i int, statInput GetStatisticsInput) {
			defer wg.Done()

			stat, err := uc.getStatistic.Do(ctx, &statInput)

			if err != nil {
				stats[i] = entities.Statistic{
					Type: statInput.Statistic.Type,
					Contract: &entities.Contract{
						Blockchain: statInput.Statistic.Contract.Blockchain,
						Address:    statInput.Statistic.Contract.Address,
						Interface:  statInput.Statistic.Contract.Interface,
					},
					Data: nil,
				}
			} else {
				stats[i] = *stat
			}

		}(i, s)
	}

	wg.Wait()

	return &stats
}
