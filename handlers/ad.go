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

type AdsHandlers struct {
	AdsUC *useCases.AdsUC
	UsersUC *useCases.UsersUC
}

func (ah *AdsHandlers) NewAd(writer http.ResponseWriter, req *http.Request) {
	var newAd models.Ad
	err := json.UnmarshalFromReader(req.Body, &newAd)
	if err != nil {
		utils.WriteError(false, "Error unmarshaling json", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
	var user models.User
	userInfo := mux.Vars(req)["id"]
	user.Id, err = strconv.Atoi(userInfo)
	if err != nil {
		utils.WriteError(false, "Error reading request", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	cookie, err := req.Cookie(utils.COOKIE_NAME)
	if err != nil {
		utils.WriteError(false, "Error finding cookie", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
	authorized, err := ah.UsersUC.CheckUser(cookie, user.Id)
	if !authorized {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Unauthorized"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
	userNotExist, err := ah.AdsUC.NewAd(&newAd, &user)
	if userNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateAnswerAdJson(writer,  utils.StatusCode("Created"), newAd)
}

func (ah *AdsHandlers) GetUserAds(writer http.ResponseWriter, req *http.Request) {
	ads := make([]models.Ad, 0)
	query := req.URL.Query()
	adType := query.Get("type")
	adTypeInt, err := strconv.Atoi(adType)
	var user models.User
	userInfo := mux.Vars(req)["id"]
	user.Id, err = strconv.Atoi(userInfo)
	if err != nil {
		utils.WriteError(false, "Error reading request", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	userNotExist, ads, err := ah.AdsUC.GetUserAds(adTypeInt, &user)
	if userNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateAnswerAdsJson(writer,  utils.StatusCode("OK"), ads)
}


func (ah *AdsHandlers) searchAds(writer http.ResponseWriter, req *http.Request, search string) {
	ads := make([]models.Ad, 0)

	ads, err := ah.AdsUC.SearchAds(search)
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateAnswerAdsJson(writer,  utils.StatusCode("OK"), ads)
}

func (ah *AdsHandlers) GetAds(writer http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	search := query.Get("search")
	if search != "" {
		ah.searchAds(writer, req, search)
		return
	}


	ads := make([]models.Ad, 0)
	adType := query.Get("type")
	sort := query.Get("sort")


	adTypeInt, err := strconv.Atoi(adType)
	if err != nil {
		utils.WriteError(false, "Error reading request", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	ads, err = ah.AdsUC.GetAds(adTypeInt, sort)
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateAnswerAdsJson(writer,  utils.StatusCode("OK"), ads)
}


func (ah *AdsHandlers) GetUserAdsNumber(writer http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	adType := query.Get("type")
	adTypeInt := 3
	var err error
	if adType != "" {
		adTypeInt, err = strconv.Atoi(adType)
		if err != nil {
			utils.WriteError(false, "Error reading request", err)
			network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
			return
		}
	}
	var user models.User
	userInfo := mux.Vars(req)["id"]
	user.Id, err = strconv.Atoi(userInfo)
	if err != nil {
		utils.WriteError(false, "Error reading request", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	userNotExist, ads, err := ah.AdsUC.GetUserAdsNumber(adTypeInt, &user)
	if userNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateErrorAnswerJson(writer, utils.StatusCode("OK"), models.CreateMessage(strconv.Itoa(ads)))
}

func (ah *AdsHandlers) GetAd(writer http.ResponseWriter, req *http.Request) {
	var ad models.Ad
	var err error
	adInfo := mux.Vars(req)["id"]
	ad.Id, err = strconv.Atoi(adInfo)
	if err != nil {
		utils.WriteError(false, "Error reading request", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	adNotExist, err := ah.AdsUC.GetAd(&ad)
	if adNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateAnswerAdJson(writer,  utils.StatusCode("OK"), ad)
}