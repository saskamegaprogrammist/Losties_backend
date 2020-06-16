package models

import "time"

//easyjson:json
type Ad struct {
	Id int `json:"id"`
	UserId int `json:"userid"`
	Type int `json:"type"`
	Comments int `json:"comments"`
	Title string `json:"title"`
	Text string `json:"text"`
	Time string `json:"time"`
	Contacts string `json:"contacts"`
	Date time.Time `json:"date"`

}

//easyjson:json
type Ads []Ad
