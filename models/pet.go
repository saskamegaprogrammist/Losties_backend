package models

type Pet struct {
	Id int `json:"id"`
	AdId int `json:"adid"`
	Name string `json:"name"`
	Animal string `json:"animal"`
	Breed string `json:"breed"`
	Color string `json:"color"`
}

