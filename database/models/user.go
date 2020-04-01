package models

import (
	"github.com/saskamegaprogrammist/Losties_backend/database"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
)

type User struct {
	Id int `json:"-"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	Nickname string `json:"nickname"`
	Phone string `json:"phone"`
	Password string `json:"-"`
}

func (user *User) SignUp() (bool, error){
	userFound := false
	db := database.GetPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return userFound, err
	}
	var userExists string
	row := transaction.QueryRow("SELECT nickname FROM users WHERE nickname = $1 OR email = $2", user.Nickname, user.Email)
	err = row.Scan(&userExists)
	if userExists != "" {
		userFound = true
		return userFound, err
	}
	row = transaction.QueryRow("INSERT INTO users (firstname, lastname, email, nickname, password) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Firstname, user.Lastname, user.Email, user.Nickname, user.Password)
	err = row.Scan(&user.Id)
	if err != nil {
		utils.WriteError(false, "Failed to scan row", err)
		errRollback := transaction.Rollback()
		if err != nil {
			utils.WriteError(true, "Error rollback", errRollback)
		}
		return userFound, err
	}
	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
	}
	return userFound, nil
}