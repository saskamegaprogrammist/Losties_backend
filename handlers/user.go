package handlers

import (
	"github.com/google/logger"
	json "github.com/mailru/easyjson"
	"github.com/saskamegaprogrammist/Losties_backend/database/models"
	"github.com/saskamegaprogrammist/Losties_backend/network"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
	"net/http"
)

func Login(writer http.ResponseWriter, req *http.Request) {
	//req.Cookie("losties_cookie")
}


func SignUp(writer http.ResponseWriter, req *http.Request) {
	//req.Cookie("losties_cookie")
	var newUser models.User
	err := json.UnmarshalFromReader(req.Body, &newUser)
	if err != nil {
		utils.WriteError(false, "Error unmarshaling json", err)
		network.CreateErrorAnswerJson(writer, 500, models.CreateError(err.Error()))
		return
	}
	usersExisting, err := newUser.SignUp()
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, 500, models.CreateError(err.Error()))
		return
	}
	if usersExisting {
		network.CreateErrorAnswerJson(writer, 409, models.CreateError("User exists"))
		return
	}
	network.CreateErrorAnswerUser(writer, 201, newUser)
}