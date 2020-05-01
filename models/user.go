package models

type User struct {
	Id int `json:"-"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Nickname string `json:"nickname"`
	Phone string `json:"phone"`
	Password string `json:"password"`
}
