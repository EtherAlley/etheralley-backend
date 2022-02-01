package entities

import "github.com/etheralley/etheralley-core-api/common"

type Statistic struct {
	Type     common.StatisticType `bson:"type" json:"type" validate:"required,oneof=SWAP"`
	Contract *Contract            `bson:"contract" json:"contract" validate:"required"`
	Data     StatisticalData      `bson:"-" json:"data"`
}

type StatisticalData interface{}

type SwapToken = struct {
	Id     string `json:"id"`
	Amount string `json:"amount"`
	Symbol string `json:"symbol"`
}

type Swap = struct {
	Id        string     `json:"id"`
	Timestamp string     `json:"timestamp"`
	AmountUSD string     `json:"amountUSD"`
	Input     *SwapToken `json:"input"`
	Output    *SwapToken `json:"output"`
}
