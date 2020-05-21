package handlers

import (
	"github.com/google/logger"
	"github.com/gorilla/mux"
	json "github.com/mailru/easyjson"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/network"
	"github.com/saskamegaprogrammist/Losties_backend/useCases"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
	"log"
	"net/http"
	"strconv"
)

type UsersHandlers struct {
	UsersUC *useCases.UsersUC
}

func (uh *UsersHandlers) Auth(writer http.ResponseWriter, req *http.Request) {
	var foundUser models.User
	cookie, err := req.Cookie(utils.COOKIE_NAME)
//	log.Println(cookie.Value)
	if err != nil {
		utils.WriteError(false, "Error finding cookie", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateError(err.Error()))
		return
	}
	loggedIn, err := uh.UsersUC.LogUser(cookie, &foundUser)
	if !loggedIn {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Unauthorized"), models.CreateError(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateError(err.Error()))
		return
	}
	network.CreateAnswerUserJson(writer,  utils.StatusCode("OK"), foundUser)
}

func (uh *UsersHandlers) SignUp(writer http.ResponseWriter, req *http.Request) {

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
	userFault, cookie, err := uh.UsersUC.Login(&newUser)
	http.SetCookie(writer, cookie)
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

func (uh *UsersHandlers) Update(writer http.ResponseWriter, req *http.Request) {
	//req.Cookie("losties_cookie")
	userIdString := mux.Vars(req)["id"]
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		utils.WriteError(false, "Error getting param", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateError(err.Error()))
		return
	}
	var updUser models.User
	err = json.UnmarshalFromReader(req.Body, &updUser)
	if err != nil {
		utils.WriteError(false, "Error unmarshaling json", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateError(err.Error()))
		return
	}
	log.Println(updUser)
	if updUser.Id == 0 {
		updUser.Id = userId
	} else if updUser.Id != userId {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateError(err.Error()))
		return
	}

	userFault, err := uh.UsersUC.UpdateUser(&updUser)
	if userFault {
		network.CreateErrorAnswerJson(writer,  utils.StatusCode("Bad Request"), models.CreateError(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer,  utils.StatusCode("Internal Server Error"), models.CreateError(err.Error()))
		return
	}

	network.CreateAnswerUserJson(writer,  utils.StatusCode("OK"), updUser)
}