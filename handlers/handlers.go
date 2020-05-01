package handlers

import (
	"github.com/saskamegaprogrammist/Losties_backend/useCases"
)

type Handlers struct {
	UsersHandlers *UsersHandlers
}

var h Handlers

func Init(usersUC *useCases.UsersUC) error {
	h.UsersHandlers = &UsersHandlers{usersUC}
	return nil
}

func GetUsersH() *UsersHandlers {
	return h.UsersHandlers
}
