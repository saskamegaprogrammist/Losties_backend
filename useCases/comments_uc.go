package useCases

import (
	"fmt"
	"github.com/saskamegaprogrammist/Losties_backend/database"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
)

type CommentsUC struct {
	AdsDB *database.AdsDB
	CommentsDB *database.CommentsDB
}


func (commentsUC *CommentsUC) NewComment(comment *models.Comment, ad *models.Ad) (bool, error) {
	err := commentsUC.AdsDB.GetAdById(ad)
	if err != nil {
		return false, err
	}
	var newAd models.Ad
	newAd.Id = ad.UserId
	err = commentsUC.AdsDB.GetAdById(&newAd)
	if err != nil {
		return false, err
	}
	if ad.Id == utils.ERROR_ID || newAd.Id == utils.ERROR_ID {
		return true, fmt.Errorf("this ad doesn't exist")
	} else {
		return false, commentsUC.CommentsDB.InsertComment(comment)
	}
}

func (commentsUC *CommentsUC) GetAdsComments(ad *models.Ad) (bool, []models.Comment, error) {
	comments := make([]models.Comment, 0)
	err := commentsUC.AdsDB.GetAdById(ad)
	if err != nil {
		return false, comments, err
	}
	if ad.Id == utils.ERROR_ID {
		return true, comments, fmt.Errorf("this ad doesn't exist")
	} else {
		comments, err = commentsUC.CommentsDB.GetCommentsByAdId(ad.Id)
		return false, comments, err
	}
}