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

func (userUC *UsersUC) UpdateInfo(user *models.User) error {
	return userUC.UsersDB.UpdateInfoUserById(user)

}

func (userUC *UsersUC) UpdateNickname(user *models.User)(bool, error) {
	id, err := userUC.UsersDB.GetUserIdByNick(user)
	if err != nil {
		return false, err
	}
	if id != utils.ERROR_ID {
		return true, fmt.Errorf("this nick is taken")
	} else {
		return false, userUC.UsersDB.UpdateUserNickById(user)
	}
}

func (userUC *UsersUC) UpdateEmail(user *models.User) (bool, error) {
	id, err := userUC.UsersDB.GetUserIdByEmail(user)
	if err != nil {
		return false, err
	}
	if id != utils.ERROR_ID {
		return true, fmt.Errorf("this email is taken")
	} else {
		return false, userUC.UsersDB.UpdateUserEmailById(user)
	}
}