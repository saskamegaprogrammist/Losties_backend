package useCases

import (
	"fmt"
	"github.com/saskamegaprogrammist/Losties_backend/database"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
)

type CoordsUC struct {
	AdsDB *database.AdsDB
	CoordsDB *database.CoordsDB
}


func (coordsUC *CoordsUC) NewCoords(coords *models.Coords, ad *models.Ad) (bool, error) {
	err := coordsUC.AdsDB.GetAdById(ad)
	if err != nil {
		return false, err
	}
	var newAd models.Ad
	newAd.Id = ad.UserId
	err = coordsUC.AdsDB.GetAdById(&newAd)
	if err != nil {
		return false, err
	}
	if ad.Id == utils.ERROR_ID || newAd.Id == utils.ERROR_ID {
		return true, fmt.Errorf("this ad doesn't exist")
	} else {
		return false, coordsUC.CoordsDB.InsertCoords(coords)
	}
}

func (coordsUC *CoordsUC) GetAdCoords(ad *models.Ad, coords *models.Coords) (bool, error) {
	err := coordsUC.AdsDB.GetAdById(ad)
	if err != nil {
		return false, err
	}
	if ad.Id == utils.ERROR_ID {
		return true, fmt.Errorf("this ad doesn't exist")
	} else {
		coords.AdId = ad.Id
		return false, coordsUC.CoordsDB.GetCoordsByAdId(coords)
	}
}

func (coordsUC *CoordsUC) GetCoords() ([]models.Coords, error) {
	return coordsUC.CoordsDB.GetCoords()
}
