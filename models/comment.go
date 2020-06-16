package models

import "time"

//easyjson:json
type Comment struct {
	Id int `json:"id"`
	UserId int `json:"userid"`
	AdId int `json:"adid"`
	Text string `json:"text"`
	Date time.Time `json:"date"`
}

//easyjson:json
type Comments []Comment
