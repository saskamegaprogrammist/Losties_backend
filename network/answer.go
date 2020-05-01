package network

import (
	json "github.com/mailru/easyjson"
	models2 "github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
	"net/http"
)

func createAnswerJson(w http.ResponseWriter, statusCode int, data []byte)  {

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)

	_, err := w.Write(data)
	if err != nil {
		utils.WriteError(false, "Error writing answer", err)
	}
}

func CreateErrorAnswerJson(writer http.ResponseWriter, statusCode int, error models2.RequestError) {
	marshalledError, err := json.Marshal(error)
	if err != nil {
		utils.WriteError(false, "Error marhalling json", err)
	}
	createAnswerJson(writer, statusCode, marshalledError)
}

func CreateAnswerUserJson(writer http.ResponseWriter, statusCode int, user models2.User) {
	marshalledUser, err := json.Marshal(user)
	if err != nil {
		utils.WriteError(false, "Error marhalling json", err)
	}
	createAnswerJson(writer, statusCode, marshalledUser)
}