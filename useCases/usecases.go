package useCases

import (
	"github.com/saskamegaprogrammist/Losties_backend/database"
)

type UseCases struct {
	UsersUC *UsersUC
	AdsUC *AdsUC
	PetsUC *PetsUC
	PicUC *PicUC
	CoordsUC *CoordsUC
	CommentsUC *CommentsUC
}

var uc UseCases

func Init(usersDB *database.UsersDB, cookiesDB *database.CookiesDB, filenameDB *database.FilenameDB, adsDB *database.AdsDB, petDB *database.PetsDB, coordsDB *database.CoordsDB, commentsDB *database.CommentsDB) error {
	uc.UsersUC = &UsersUC{usersDB, cookiesDB}
	uc.AdsUC = &AdsUC{usersDB, adsDB}
	uc.PetsUC = &PetsUC{ adsDB, petDB}
	uc.PicUC = &PicUC{ adsDB, usersDB, filenameDB}
	uc.CoordsUC = &CoordsUC{adsDB, coordsDB}
	uc.CommentsUC = &CommentsUC{adsDB, commentsDB}
	return nil
}

func GetUsersUC() *UsersUC {
	return uc.UsersUC
}

func GetCoordsUC() *CoordsUC {
	return uc.CoordsUC
}
func GetCommentsUC() *CommentsUC {
	return uc.CommentsUC
}

func GetAdsUC() *AdsUC {
	return uc.AdsUC
}

func GetPetsUC() *PetsUC {
	return uc.PetsUC
}

func GetPicUC() *PicUC {
	return uc.PicUC
}