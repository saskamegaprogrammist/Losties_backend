package models

import (
	"github.com/google/logger"
	"github.com/saskamegaprogrammist/Losties_backend/database"
)

type User struct {
	Id int `json:"-"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Nickname string `json:"nickname"`
	Phone string `json:"fullname"`
	Password string `json:"-"`
}

func (user *User) CreateUser() {
	db := database.GetPool()
	_, err := db.Begin()
	if err != nil {
		logger.Errorf("Failed to start transaction %v", err)
	}
}