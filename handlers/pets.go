package handlers

import (
	"github.com/google/logger"
	"github.com/gorilla/mux"
	json "github.com/mailru/easyjson"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/network"
	"github.com/saskamegaprogrammist/Losties_backend/useCases"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
	"net/http"
	"strconv"
)


type PetsHandlers struct {
	PetsUC *useCases.PetsUC
}

func (ph *PetsHandlers) NewPet(writer http.ResponseWriter, req *http.Request) {
	var newPet models.Pet
	err := json.UnmarshalFromReader(req.Body, &newPet)
	if err != nil {
		utils.WriteError(false, "Error unmarshaling json", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
	var ad models.Ad
	adInfo := mux.Vars(req)["id"]
	ad.Id, err = strconv.Atoi(adInfo)
	if err != nil {
		utils.WriteError(false, "Error reading request", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	adNotExist, err := ph.PetsUC.NewPet(&newPet, &ad)
	if adNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateAnswerPetJson(writer,  utils.StatusCode("Created"), newPet)
}

func (ph *PetsHandlers) GetAdPet(writer http.ResponseWriter, req *http.Request) {
	var pet models.Pet
	var ad models.Ad
	adInfo := mux.Vars(req)["id"]
	var err error
	ad.Id, err = strconv.Atoi(adInfo)
	if err != nil {
		utils.WriteError(false, "Error reading request", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	adNotExist, err := ph.PetsUC.GetAdPet(&ad, &pet)
	if adNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateAnswerPetJson(writer,  utils.StatusCode("OK"), pet)
}

