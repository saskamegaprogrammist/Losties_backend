package useCases

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/saskamegaprogrammist/Losties_backend/database"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
	"log"
	"net/http"
	"time"
)

type UsersUC struct {
	UsersDB *database.UsersDB
	CookiesDB *database.CookiesDB
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

func (userUC *UsersUC) Login(user *models.User) (bool, *http.Cookie, error) {
	err := userUC.UsersDB.GetUserByEmailAndPassword(user)
	if user.Id == utils.ERROR_ID {
		return true, nil, fmt.Errorf("email or password is wrong")
	} else {
		if err != nil {
			return false, nil,err
		}
		cookie, err := userUC.SetCookie(user)
		return false, &cookie, err
	}
}

func (userUC *UsersUC) LogUser(cookie *http.Cookie, user *models.User) (bool, error) {
	cookieInfo, err := userUC.CookiesDB.GetUserIdByCookie(cookie.Value)
	log.Println(cookieInfo)
	if cookieInfo.User == utils.ERROR_ID {
		return false, fmt.Errorf("not logged in")
	}
	if err != nil {
		return true, err
	}
	user.Id = cookieInfo.User
	err = userUC.UsersDB.GetUserById(user)
	return true, err
}

func (userUC *UsersUC) SetCookie(user *models.User) (http.Cookie, error) {
	token := uuid.New()
	sessionExpiration := time.Now().Add(365 * 24 * time.Hour)
	cookie := http.Cookie{Name: utils.COOKIE_NAME, Value: token.String(), Expires: sessionExpiration}
	var cookieInfo models.CookieInfo
	cookieInfo.User = user.Id
	cookieInfo.Cookie = cookie.Value
	return cookie, userUC.CookiesDB.InsertCookie(&cookieInfo)
}

func (userUC *UsersUC) UpdateUser(newUser *models.User) (bool, error) {
	var oldUser models.User
	oldUser.Id = newUser.Id
	err := userUC.UsersDB.GetUserById(&oldUser)
	if err != nil {
		return false, err
	}
	if oldUser.Nickname != newUser.Nickname {
		nickExists, err := userUC.UpdateNickname(newUser)
		if nickExists || err != nil {
			return nickExists, err
		}
	}
	if oldUser.Email != newUser.Email {
		emailExists, err := userUC.UpdateEmail(newUser)
		if emailExists || err != nil {
			return emailExists, err
		}
	}
	return false, userUC.UsersDB.UpdateInfoUserById(newUser)
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