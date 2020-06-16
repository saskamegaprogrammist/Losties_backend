package useCases

import (
	"fmt"
	"github.com/saskamegaprogrammist/Losties_backend/database"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
)

type PetsUC struct {
	AdsDB *database.AdsDB
	PetsDB *database.PetsDB
}


func (petsUC *PetsUC) NewPet(pet *models.Pet, ad *models.Ad) (bool, error) {
	err := petsUC.AdsDB.GetAdById(ad)
	if err != nil {
		return false, err
	}
	var newAd models.Ad
	newAd.Id = ad.UserId
	err = petsUC.AdsDB.GetAdById(&newAd)
	if err != nil {
		return false, err
	}
	if ad.Id == utils.ERROR_ID || newAd.Id == utils.ERROR_ID {
		return true, fmt.Errorf("this ad doesn't exist")
	} else {
		return false, petsUC.PetsDB.InsertPet(pet)
	}
}

func (petsUC *PetsUC) GetAdPet(ad *models.Ad, pet *models.Pet) (bool, error) {
	err := petsUC.AdsDB.GetAdById(ad)
	if err != nil {
		return false, err
	}
	if ad.Id == utils.ERROR_ID {
		return true, fmt.Errorf("this ad doesn't exist")
	} else {
		pet.AdId = ad.Id
		return false, petsUC.PetsDB.GetPetByAdId(pet)
	}
}
