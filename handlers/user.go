package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/saskamegaprogrammist/Losties_backend/database/models"
	"net/http"
)

func Login(writer http.ResponseWriter, req *http.Request) {
	//req.Cookie("losties_cookie")
}


func SignUp(writer http.ResponseWriter, req *http.Request) {
	//req.Cookie("losties_cookie")
	var newUser models.User
	err := json.NewDecoder(req.Body).Decode(&newUser)
	if err != nil {
		//log.Println(err)
		utils.CreateAnswer(writer, 500, models.CreateError("cannot decode json"))
		return
	}
	userNickname := mux.Vars(req)["nickname"]
	newUser.Nickname = userNickname
	usersExisting, err := newUser.CreateUser()
	if err != nil {
		//log.Println(err)
		utils.CreateAnswer(writer, 500, models.CreateError("internal error"))
		return
	}
	if usersExisting != nil {
		utils.CreateAnswer(writer, 409, usersExisting)
		return
	}
	utils.CreateAnswer(writer, 201, newUser)
}