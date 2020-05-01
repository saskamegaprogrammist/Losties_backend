package handlers

import (
	"github.com/google/logger"
	json "github.com/mailru/easyjson"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/network"
	"github.com/saskamegaprogrammist/Losties_backend/useCases"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
	"net/http"
)

type UsersHandlers struct {
	UsersUC *useCases.UsersUC
}



func (uh *UsersHandlers) SignUp(writer http.ResponseWriter, req *http.Request) {
	//req.Cookie("losties_cookie")
	var newUser models.User
	err := json.UnmarshalFromReader(req.Body, &newUser)
	if err != nil {
		utils.WriteError(false, "Error unmarshaling json", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateError(err.Error()))
		return
	}
	usersExisting, err := uh.UsersUC.SignUp(&newUser)
	if usersExisting {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Conflict"), models.CreateError(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateError(err.Error()))
		return
	}

	network.CreateAnswerUserJson(writer,  utils.StatusCode("Created"), newUser)
}

func (uh *UsersHandlers) Login(writer http.ResponseWriter, req *http.Request) {
	//req.Cookie("losties_cookie")
	var newUser models.User
	err := json.UnmarshalFromReader(req.Body, &newUser)
	if err != nil {
		utils.WriteError(false, "Error unmarshaling json", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateError(err.Error()))
		return
	}
	userFault, err := uh.UsersUC.Login(&newUser)
	if userFault {
		network.CreateErrorAnswerJson(writer,  utils.StatusCode("Bad Request"), models.CreateError(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer,  utils.StatusCode("Internal Server Error"), models.CreateError(err.Error()))
		return
	}

	network.CreateAnswerUserJson(writer,  utils.StatusCode("OK"), newUser)
}