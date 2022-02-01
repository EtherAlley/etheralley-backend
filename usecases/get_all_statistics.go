package usecases

import (
	"context"
	"sync"

	"github.com/etheralley/etheralley-core-api/common"
	"github.com/etheralley/etheralley-core-api/entities"
)

func NewGetAllStatisticsUseCase(logger *common.Logger, getStatistic IGetStatisticUseCase) IGetAllStatisticsUseCase {
	return GetAllStatistics(logger, getStatistic)
}

func GetAllStatistics(logger *common.Logger, getStatistic IGetStatisticUseCase) IGetAllStatisticsUseCase {
	return func(ctx context.Context, input *GetAllStatisticsInput) *[]entities.Statistic {
		var wg sync.WaitGroup

		stats := make([]*entities.Statistic, len(*input.Stats))

		for i, s := range *input.Stats {
			wg.Add(1)

			go func(
				i int,
				s StatisticInput) {
				defer wg.Done()

				stat, err := getStatistic(ctx, input.Address, s.Contract, s.Type)

				if err != nil {
					logger.Errf(err, "invalid swaps contract: type %v address %v chain %v interface %v", s.Type, s.Contract.Address, s.Contract.Blockchain, s.Contract.Interface)
					return
				}

				stats[i] = stat
			}(i, s)
		}

		wg.Wait()

		trimmedStats := []entities.Statistic{}
		for _, stat := range stats {
			if stat != nil {
				trimmedStats = append(trimmedStats, *stat)
			}
		}

		return &trimmedStats
	}
}
