package entities

import (
	"gorm.io/datatypes"
)

type Element struct {
	Id             string         `gorm:"primaryKey" json:"id"`
	ProfileAddress string         `gorm:"primaryKey" json:"-"`
	Type           string         `json:"type"`
	PositionX      uint           `json:"positionX"`
	PositionY      uint           `json:"positionY"`
	Data           datatypes.JSON `json:"data"`
}
