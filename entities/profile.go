package entities

type Element struct {
	Id        string                 `bson:"id" json:"id"`
	Type      string                 `bson:"type" json:"type"`
	PositionX uint                   `bson:"position_x" json:"position_x"`
	PositionY uint                   `bson:"position_y" json:"position_y"`
	Data      map[string]interface{} `bson:"data" json:"data"`
}

type Profile struct {
	Address  string    `bson:"_id" json:"address"`
	Elements []Element `bson:"elements" json:"elements"`
}
