package database

import (
	"database/sql"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
)

type UsersDB struct {
}


func (usersDB *UsersDB) UpdateUser(user *models.User) error {
	//db := getPool()
	//transaction, err := db.Begin()
	//if err != nil {
	//	utils.WriteError(false, "Failed to start transaction", err)
	//	return err
	//}
	//

}

func (usersDB *UsersDB) GetUserIdByNickAndEmail(user *models.User) (int, error) {
	userExistsId := utils.ERROR_ID
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return userExistsId, err
	}


	row := transaction.QueryRow("SELECT id FROM users WHERE nickname = $1 OR email = $2", user.Nickname, user.Email)
	row.Scan(&userExistsId)

	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return userExistsId, err
	}
	return userExistsId, nil
}

func (usersDB *UsersDB) GetUserByEmailAndPassword(user *models.User) error {
	db := getPool()
	transaction, err := db.Begin()if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}

	var userPhone sql.NullString
	user.Id = utils.ERROR_ID
	rows := transaction.QueryRow("SELECT * FROM users WHERE email = $1 and password = $2", user.Email, user.Password)
	err = rows.Scan(&user.Id, &user.Firstname,&user.Lastname, &user.Email, &user.Nickname, &userPhone, &user.Password)
	if err != nil {
		utils.WriteError(false, "Failed to scan row", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			utils.WriteError(true, "Error rollback", errRollback)
		}
		return err
	}
	if userPhone.Valid {
		user.Phone = userPhone.String
	} else {
		user.Phone = ""
	}

	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return err
	}
	return nil
}


func (usersDB *UsersDB) InsertUser(user *models.User) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}

	row := transaction.QueryRow("INSERT INTO users (firstname, lastname, email, nickname, password) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Firstname, user.Lastname, user.Email, user.Nickname, user.Password)
	err = row.Scan(&user.Id)
	if err != nil {
		utils.WriteError(false, "Failed to scan row", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			utils.WriteError(true, "Error rollback", errRollback)
		}
		return err
	}

	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return err
	}
	return nil
}
