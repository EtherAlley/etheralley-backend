package entities

type Statistic struct {
	Contract *Contract       `bson:"contract" json:"contract" validate:"required"`
	Data     StatisticalData `bson:"-" json:"data"`
}

type StatisticalData interface{}
