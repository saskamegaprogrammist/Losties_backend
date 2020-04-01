package network

import (
	json "github.com/mailru/easyjson"
	"github.com/saskamegaprogrammist/Losties_backend/database/models"
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

func CreateErrorAnswerJson(writer http.ResponseWriter, statusCode int, error models.RequestError) {
	marshalledError, err := json.Marshal(error)
	if err != nil {
		utils.WriteError(false, "Error marhalling json", err)
	}
	createAnswerJson(writer, statusCode, marshalledError)
}

func CreateErrorAnswerUser(writer http.ResponseWriter, statusCode int, user models.User) {
	marshalledUser, err := json.Marshal(user)
	if err != nil {
		utils.WriteError(false, "Error marhalling json", err)
	}
	createAnswerJson(writer, statusCode, marshalledUser)
}