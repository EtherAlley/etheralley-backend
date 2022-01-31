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
	return func(ctx context.Context, address string, contract *[]entities.Contract) *[]entities.Statistic {
		var wg sync.WaitGroup

		stats := make([]*entities.Statistic, len(*contract))

		for i, contract := range *contract {
			wg.Add(1)

			go func(i int, c entities.Contract) {
				defer wg.Done()

				stat, err := getStatistic(ctx, address, &c)

				if err != nil {
					logger.Errf(err, "invalid swaps contract: address %v chain %v interface %v", c.Address, c.Blockchain, c.Interface)
					return
				}

				stats[i] = stat
			}(i, contract)
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
