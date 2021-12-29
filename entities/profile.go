package entities

type Profile struct {
	Address  string    `bson:"_id" json:"-"`
	Elements []Element `bson:"elements" json:"elements"`
}

type Element struct {
	Id       string                 `bson:"id" json:"id"`
	Type     string                 `bson:"type" json:"type"`
	Position XYPosition             `bson:"position" json:"position"`
	Data     map[string]interface{} `bson:"data" json:"data"`
}

type XYPosition struct {
	X int `bson:"x" json:"x"`
	Y int `bson:"y" json:"y"`
}
