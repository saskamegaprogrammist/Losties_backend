package useCases

import (
	"github.com/saskamegaprogrammist/Losties_backend/database"
)

type UseCases struct {
	UsersUC *UsersUC
}

var uc UseCases

func Init(usersDB *database.UsersDB) error {
	uc.UsersUC = &UsersUC{usersDB}
	return nil
}

func GetUsersUC() *UsersUC {
	return uc.UsersUC
}