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

type CommentsHandlers struct {
	CommentsUC *useCases.CommentsUC
}

func (ch *CommentsHandlers) NewComment(writer http.ResponseWriter, req *http.Request) {
	var newComment models.Comment
	err := json.UnmarshalFromReader(req.Body, &newComment)
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
	adNotExist, err := ch.CommentsUC.NewComment(&newComment, &ad)
	if adNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateAnswerCommentJson(writer,  utils.StatusCode("Created"), newComment)
}

func (ch *CommentsHandlers) GetAdComments(writer http.ResponseWriter, req *http.Request) {
	comments := make([]models.Comment, 0)
	var ad models.Ad
	var err error
	adInfo := mux.Vars(req)["id"]
	ad.Id, err = strconv.Atoi(adInfo)
	if err != nil {
		utils.WriteError(false, "Error reading request", err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	adNotExist, comments, err := ch.CommentsUC.GetAdsComments(&ad)
	if adNotExist {
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Bad Request"), models.CreateMessage(err.Error()))
		return
	}
	if err != nil {
		logger.Error(err)
		network.CreateErrorAnswerJson(writer, utils.StatusCode("Internal Server Error"), models.CreateMessage(err.Error()))
		return
	}

	network.CreateAnswerCommentsJson(writer,  utils.StatusCode("OK"), comments)
}
