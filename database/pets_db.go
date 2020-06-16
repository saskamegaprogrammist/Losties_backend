package database

import (
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
)

type PetsDB struct {
}

func (petsDB *PetsDB) InsertPet(pet *models.Pet) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}

	row := transaction.QueryRow("INSERT INTO pets (adid, name, animal, breed, color) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		pet.AdId, pet.Name, pet.Animal, pet.Breed, pet.Color)
	err = row.Scan(&pet.Id)
	if err != nil {
		utils.WriteError(false, "Failed to scan row", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			utils.WriteError(true, "Error rollback", errRollback)
		}
		return err
	}

	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return err
	}
	return nil
}

func (petsDB *PetsDB) GetPetByAdId(pet *models.Pet) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}

	rows := transaction.QueryRow("SELECT * FROM pets WHERE adid = $1", pet.AdId)
	err = rows.Scan(&pet.Id, &pet.AdId,&pet.Name, &pet.Animal, &pet.Breed, &pet.Color)
	if err != nil {
		utils.WriteError(false, "Failed to scan row", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			utils.WriteError(true, "Error rollback", errRollback)
		}
		return err
	}
	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return err
	}
	return nil
}
