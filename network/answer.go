package network

import (
	json "github.com/mailru/easyjson"
	models2 "github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
	"net/http"
)

func createAnswerJson(w http.ResponseWriter, statusCode int, data []byte)  {
	w.WriteHeader(statusCode)
	w.Header().Set("content-type", "application/json")
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

func CreateAnswerUserPublicJson(writer http.ResponseWriter, statusCode int, user models2.UserPublic) {
	marshalledUser, err := json.Marshal(user)
	if err != nil {
		utils.WriteError(false, "Error marhalling json", err)
	}
	createAnswerJson(writer, statusCode, marshalledUser)
}

func CreateAnswerAdJson(writer http.ResponseWriter, statusCode int, ad models2.Ad) {
	marshalledAd, err := json.Marshal(ad)
	if err != nil {
		utils.WriteError(false, "Error marhalling json", err)
	}
	createAnswerJson(writer, statusCode, marshalledAd)
}

func CreateAnswerPetJson(writer http.ResponseWriter, statusCode int, ad models2.Pet) {
	marshalledPet, err := json.Marshal(ad)
	if err != nil {
		utils.WriteError(false, "Error marhalling json", err)
	}
	createAnswerJson(writer, statusCode, marshalledPet)
}

func CreateAnswerCoordsJson(writer http.ResponseWriter, statusCode int, ad models2.Coords) {
	marshalledC, err := json.Marshal(ad)
	if err != nil {
		utils.WriteError(false, "Error marhalling json", err)
	}
	createAnswerJson(writer, statusCode, marshalledC)
}

func CreateAnswerCommentJson(writer http.ResponseWriter, statusCode int, comment models2.Comment) {
	marshalledC, err := json.Marshal(comment)
	if err != nil {
		utils.WriteError(false, "Error marhalling json", err)
	}
	createAnswerJson(writer, statusCode, marshalledC)
}

func CreateAnswerCommentsJson(writer http.ResponseWriter, statusCode int, comments models2.Comments) {
	marshalledC, err := json.Marshal(comments)
	if err != nil {
		utils.WriteError(false, "Error marhalling json", err)
	}
	createAnswerJson(writer, statusCode, marshalledC)
}

func CreateAnswerCoordsAllJson(writer http.ResponseWriter, statusCode int, coordsAll models2.CoordsAll) {
	marshalledCA, err := json.Marshal(coordsAll)
	if err != nil {
		utils.WriteError(false, "Error marhalling json", err)
	}
	createAnswerJson(writer, statusCode, marshalledCA)
}

func CreateAnswerAdsJson(writer http.ResponseWriter, statusCode int, ads models2.Ads) {
	marshalledAd, err := json.Marshal(ads)
	if err != nil {
		utils.WriteError(false, "Error marhalling json", err)
	}
	createAnswerJson(writer, statusCode, marshalledAd)
}
