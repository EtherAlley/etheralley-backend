package entities

import "github.com/etheralley/etheralley-core-api/common"

type Statistic struct {
	Type     common.StatisticType
	Contract *Contract
	Data     interface{}
}
