package entities

import "github.com/etheralley/etheralley-backend/common"

type Statistic struct {
	Type     common.StatisticType
	Contract *Contract
	Data     interface{}
}
