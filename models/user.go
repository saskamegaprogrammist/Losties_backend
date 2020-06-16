package models

//easyjson:json
type User struct {
	Id int `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Nickname string `json:"nickname"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}

//easyjson:json
type UserPublic struct {
	Id int `json:"id"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Nickname string `json:"nickname"`
	Phone string `json:"-"`
	Password string `json:"-"`
}


