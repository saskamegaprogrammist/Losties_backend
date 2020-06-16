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

type CoordsHandlers struct {
	CoordsUC *useCases.CoordsUC
}


func (ch *CoordsHandlers) NewCoords(writer http.ResponseWriter, req *http.Request) {
	var newCoords models.Coords
	err := json.UnmarshalFromReader(req.Body, &newCoords)
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
	adNotExist, err := ch.CoordsUC.NewCoords(&newCoords, &ad)
	if adNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateAnswerCoordsJson(writer,  utils.StatusCode("Created"), newCoords)
}

func (ch *CoordsHandlers) GetCoords(writer http.ResponseWriter, req *http.Request) {
	coords := make([]models.Coords, 0)

	coords, err := ch.CoordsUC.GetCoords()
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateAnswerCoordsAllJson(writer,  utils.StatusCode("OK"), coords)
}

func (ch *CoordsHandlers) GetAdCoords(writer http.ResponseWriter, req *http.Request) {
	var coords models.Coords
	var ad models.Ad
	adInfo := mux.Vars(req)["id"]
	var err error
	ad.Id, err = strconv.Atoi(adInfo)
	if err != nil {
		utils.WriteError(false, "Error reading request", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	adNotExist, err := ch.CoordsUC.GetAdCoords(&ad, &coords)
	if adNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateAnswerCoordsJson(writer,  utils.StatusCode("OK"), coords)
}
