package useCases

import (
	"fmt"
	"github.com/saskamegaprogrammist/Losties_backend/database"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
)

type AdsUC struct {
	UsersDB *database.UsersDB
	AdsDB *database.AdsDB
}


func (adsUC *AdsUC) NewAd(ad *models.Ad, user *models.User) (bool, error) {
	err := adsUC.UsersDB.GetUserById(user)
	if err != nil {
		return false, err
	}
	var newUser models.User
	newUser.Id = ad.UserId
	err = adsUC.UsersDB.GetUserById(&newUser)
	if err != nil {
		return false, err
	}
	if user.Id == utils.ERROR_ID || newUser.Id == utils.ERROR_ID {
		return true, fmt.Errorf("this user doesn't exist")
	} else {
		return false, adsUC.AdsDB.InsertAd(ad)
	}
}

func (adsUC *AdsUC) GetUserAds(adType int, user *models.User) (bool, []models.Ad, error) {
	ads := make([]models.Ad, 0)
	err := adsUC.UsersDB.GetUserById(user)
	if err != nil {
		return false, ads, err
	}
	if user.Id == utils.ERROR_ID {
		return true, ads, fmt.Errorf("this user doesn't exist")
	} else {
		ads, err = adsUC.AdsDB.GetAdsByUserId(user.Id, adType)
		return false, ads, err
	}
}

func (adsUC *AdsUC) GetUserAdsNumber(adType int, user *models.User) (bool, int, error) {
	ads := 0
	err := adsUC.UsersDB.GetUserById(user)
	if err != nil {
		return false, ads, err
	}
	if user.Id == utils.ERROR_ID {
		return true, ads, fmt.Errorf("this user doesn't exist")
	} else {
		ads, err = adsUC.AdsDB.GetAdsNumberByUserId(user.Id, adType)
		return false, ads, err
	}
}

func (adsUC *AdsUC) GetAds(adType int, sort string) ([]models.Ad, error) {
	ads := make([]models.Ad, 0)
	var err error
	if sort == "" {
		ads, err = adsUC.AdsDB.GetAds(adType)
		return ads, err
	} else {
		ads, err = adsUC.AdsDB.GetAdsSorted (adType, sort)
		return ads, err
	}
}

func (adsUC *AdsUC) SearchAds(search string) ([]models.Ad, error) {
	ads := make([]models.Ad, 0)
	var err error
	ads, err = adsUC.AdsDB.SearchAds (search)
	return ads, err
}

func (adsUC *AdsUC) GetAd(ad *models.Ad) (bool, error) {
	err := adsUC.AdsDB.GetAdById(ad)
	if err != nil {
		return false, err
	}
	if ad.Id == utils.ERROR_ID {
		return true, fmt.Errorf("this ad doesn't exist")
	} else {
		return false, nil
	}
}