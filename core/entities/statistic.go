package entities

import "github.com/etheralley/etheralley-apis/common"

type Statistic struct {
	Type     common.StatisticType
	Contract *Contract
	Data     interface{}
}
