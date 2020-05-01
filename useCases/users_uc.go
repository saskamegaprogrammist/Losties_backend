package useCases

import (
	"fmt"
	"github.com/saskamegaprogrammist/Losties_backend/database"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
)

type UsersUC struct {
	UsersDB *database.UsersDB
}

func (userUC *UsersUC) SignUp(user *models.User) (bool, error) {
	id, err := userUC.UsersDB.GetUserIdByNickAndEmail(user)
	if err != nil {
		return false, err
	}
	if id != utils.ERROR_ID {
		return true, fmt.Errorf("this user exists")
	} else {
		return false, userUC.UsersDB.InsertUser(user)
	}
}

func (userUC *UsersUC) Login(user *models.User) (bool, error) {
	err := userUC.UsersDB.GetUserByEmailAndPassword(user)
	if user.Id == utils.ERROR_ID {
		return true, fmt.Errorf("email or password is wrong")
	} else {
		return false, err
	}
}