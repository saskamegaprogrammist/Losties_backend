package handlers

import (
	"bufio"
	"github.com/google/logger"
	"github.com/gorilla/mux"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/network"
	"github.com/saskamegaprogrammist/Losties_backend/useCases"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
	"net/http"
	"strconv"
)

type PicHandlers struct {
	PicUC *useCases.PicUC
	UsersUC *useCases.UsersUC
}

func (ph *PicHandlers) NewAdPic(writer http.ResponseWriter, req *http.Request) {
	var ad models.Ad
	var err error
	adInfo := mux.Vars(req)["id"]
	ad.Id, err = strconv.Atoi(adInfo)
	if err != nil {
		utils.WriteError(false, "Error reading request", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	err = req.ParseMultipartForm(32 << 20)
	if err != nil {
		utils.WriteError(false, "Error parsing picture", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
	file, _, err := req.FormFile("adpic")
	if err != nil {
		utils.WriteError(false, "Error parsing picture", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
	defer file.Close()

	adNotExist, err := ph.PicUC.NewAdPic(&ad, file, adInfo)
	if adNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateErrorAnswerJson(writer, utils.StatusCode("Created"), models.CreateMessage("OK"))
}

func (ph *PicHandlers) AdPicGet(writer http.ResponseWriter, req *http.Request) {
	var ad models.Ad
	var err error
	adInfo := mux.Vars(req)["id"]
	ad.Id, err = strconv.Atoi(adInfo)
	if err != nil {
		utils.WriteError(false, "Error reading request", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	adNotExist, file, err := ph.PicUC.GetAdPic(&ad)
	if adNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()

	bytes := make([]byte, size)
	_, err = reader.Read(bytes)

	writer.Header().Set("content-type", "multipart/form-data;boundary=1")

	_, err = writer.Write(bytes)
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
}

func (ph *PicHandlers) NewUserPic(writer http.ResponseWriter, req *http.Request) {
	var user models.User
	var err error
	adInfo := mux.Vars(req)["id"]
	user.Id, err = strconv.Atoi(adInfo)
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
	authorized, err := ph.UsersUC.CheckUser(cookie, user.Id)
	if !authorized {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Unauthorized"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	err = req.ParseMultipartForm(32 << 20)
	if err != nil {
		utils.WriteError(false, "Error parsing picture", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
	file, _, err := req.FormFile("userpic")
	if err != nil {
		utils.WriteError(false, "Error parsing picture", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
	defer file.Close()

	userNotExist, err := ph.PicUC.NewUserPic(&user, file, adInfo)
	if userNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateErrorAnswerJson(writer, utils.StatusCode("Created"), models.CreateMessage("OK"))
}

func (ph *PicHandlers) UserPicGet(writer http.ResponseWriter, req *http.Request) {
	var user models.User
	var err error
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
	authorized, err := ph.UsersUC.CheckUser(cookie, user.Id)
	if !authorized {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Unauthorized"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	userNotExist, file, err := ph.PicUC.GetUserPic(&user)
	if userNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()

	bytes := make([]byte, size)
	_, err = reader.Read(bytes)

	writer.Header().Set("content-type", "multipart/form-data;boundary=1")

	_, err = writer.Write(bytes)
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}
}
