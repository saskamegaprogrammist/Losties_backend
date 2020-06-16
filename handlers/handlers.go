package handlers

import (
	"github.com/saskamegaprogrammist/Losties_backend/useCases"
)

type Handlers struct {
	UsersHandlers *UsersHandlers
	AdsHandlers *AdsHandlers
	PetsHandlers *PetsHandlers
	PicHandlers *PicHandlers
	CoordsHandlers *CoordsHandlers
	CommentsHandlers *CommentsHandlers
}

var h Handlers

func Init(usersUC *useCases.UsersUC, adsUC *useCases.AdsUC, petsUC *useCases.PetsUC, picUC *useCases.PicUC, coordsUC *useCases.CoordsUC, commentsUC *useCases.CommentsUC) error {
	h.UsersHandlers = &UsersHandlers{usersUC}
	h.AdsHandlers = &AdsHandlers{adsUC, usersUC}
	h.PetsHandlers = &PetsHandlers{petsUC}
	h.PicHandlers = &PicHandlers{picUC, usersUC}
	h.CoordsHandlers = &CoordsHandlers{coordsUC}
	h.CommentsHandlers = &CommentsHandlers{commentsUC}
	return nil
}

func GetUsersH() *UsersHandlers {
	return h.UsersHandlers
}

func GetCoordsH() *CoordsHandlers {
	return h.CoordsHandlers
}

func GetCommentsH() *CommentsHandlers {
	return h.CommentsHandlers
}

func GetAdsH() *AdsHandlers {
	return h.AdsHandlers
}

func GetPetsH() *PetsHandlers {
	return h.PetsHandlers
}

func GetPicH() *PicHandlers {
	return h.PicHandlers
}