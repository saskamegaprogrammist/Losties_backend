package models

//easyjson:json
type Coords struct {
	Id int `json:"id"`
	AdId int `json:"adid"`
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

//easyjson:json
type CoordsAll []Coords