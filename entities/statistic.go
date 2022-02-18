package entities

import "github.com/etheralley/etheralley-core-api/common"

type Statistic struct {
	Type     common.StatisticType
	Contract *Contract
	Data     StatisticalData
}

type StatisticalData interface{}

type SwapToken = struct {
	Id     string
	Amount string
	Symbol string
}

type Swap = struct {
	Id        string
	Timestamp string
	AmountUSD string
	Input     *SwapToken
	Output    *SwapToken
}
