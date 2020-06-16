package database

import (
	"database/sql"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
)

type UsersDB struct {
}

func (usersDB *UsersDB) UpdateUserNickById(user *models.User) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}
	var userPhone sql.NullString
	rows := transaction.QueryRow("UPDATE users SET nickname=$1 WHERE id=$2 returning firstname, lastname, email, nickname, phone", user.Nickname, user.Id)
	err = rows.Scan(&user.Firstname,&user.Lastname, &user.Email,  &user.Nickname, &userPhone)
	if err != nil {
		utils.WriteError(false, "Failed to scan row", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			utils.WriteError(true, "Error rollback", errRollback)
		}
		return err
	}
	user.Phone = checkNullString(userPhone)

	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return err
	}
	return nil
}

func (usersDB *UsersDB) UpdateUserEmailById(user *models.User) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}
	var userPhone sql.NullString
	rows := transaction.QueryRow("UPDATE users SET email=$1 WHERE id=$2 returning firstname, lastname, email, nickname, phone", user.Email, user.Id)
	err = rows.Scan(&user.Firstname,&user.Lastname, &user.Email,  &user.Nickname, &userPhone)
	if err != nil {
		utils.WriteError(false, "Failed to scan row", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			utils.WriteError(true, "Error rollback", errRollback)
		}
		return err
	}
	user.Phone = checkNullString(userPhone)

	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return err
	}
	return nil
}

func (usersDB *UsersDB) UpdateInfoUserById(user *models.User) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}
	var userPhone sql.NullString
	rows := transaction.QueryRow("UPDATE users SET firstname = coalesce(nullif($1, ''), firstname),  lastname = coalesce(nullif($2, ''), lastname) , phone = coalesce(nullif($3, ''), phone) WHERE id=$4 returning firstname, lastname, email, nickname, phone", user.Firstname, user.Lastname, user.Phone, user.Id)
	err = rows.Scan(&user.Firstname,&user.Lastname, &user.Email,  &user.Nickname, &userPhone)
	if err != nil {
		utils.WriteError(false, "Failed to scan row", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			utils.WriteError(true, "Error rollback", errRollback)
		}
		return err
	}
	user.Phone = checkNullString(userPhone)

	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return err
	}
	return nil

}

func (usersDB *UsersDB) GetUserIdByEmail(user *models.User) (int, error) {
	userExistsId := utils.ERROR_ID
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return userExistsId, err
	}


	row := transaction.QueryRow("SELECT id FROM users WHERE email = $1", user.Email)
	row.Scan(&userExistsId)

	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return userExistsId, err
	}
	return userExistsId, nil
}

func (usersDB *UsersDB) GetUserIdByNick(user *models.User) (int, error) {
	userExistsId := utils.ERROR_ID
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return userExistsId, err
	}


	row := transaction.QueryRow("SELECT id FROM users WHERE nickname = $1", user.Nickname)
	row.Scan(&userExistsId)

	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return userExistsId, err
	}
	return userExistsId, nil
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

func (usersDB *UsersDB) GetUserById(user *models.User) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}

	var userPhone sql.NullString
	rows := transaction.QueryRow("SELECT * FROM users WHERE id = $1", user.Id)
	err = rows.Scan(&user.Id, &user.Firstname,&user.Lastname, &user.Email, &user.Nickname, &userPhone, &user.Password)
	if err != nil {
		utils.WriteError(false, "Failed to scan row", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			utils.WriteError(true, "Error rollback", errRollback)
		}
		return err
	}
	user.Phone = checkNullString(userPhone)

	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return err
	}
	return nil
}

func (usersDB *UsersDB) GetUserPublicById(user *models.UserPublic) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}

	rows := transaction.QueryRow("SELECT id, firstname, lastname, email, nickname FROM users WHERE id = $1", user.Id)
	err = rows.Scan(&user.Id, &user.Firstname,&user.Lastname, &user.Email, &user.Nickname)
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

func (usersDB *UsersDB) GetUserByEmailAndPassword(user *models.User) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
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
	user.Phone = checkNullString(userPhone)

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

	row := transaction.QueryRow("INSERT INTO users (firstname, lastname, email, nickname, password, phone) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		user.Firstname, user.Lastname, user.Email, user.Nickname, user.Password, user.Phone)
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

func checkNullString(sqlStr sql.NullString) string {
	if sqlStr.Valid {
		return sqlStr.String
	} else {
		return ""
	}
}