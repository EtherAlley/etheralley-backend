package entities

type Profile struct {
	Address  string    `gorm:"primaryKey" json:"address"`
	Elements []Element `gorm:"foreignKey:ProfileAddress" json:"elements"`
}
