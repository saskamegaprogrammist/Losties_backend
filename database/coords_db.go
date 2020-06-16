package database

import (
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
)

type CoordsDB struct {
}

func (coordsDB *CoordsDB) InsertCoords (coords *models.Coords) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}

	row := transaction.QueryRow("INSERT INTO coords (adid, x, y) VALUES ($1, $2, $3) RETURNING id",
		coords.AdId, coords.X, coords.Y)
	err = row.Scan(&coords.Id)
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

func (coordsDB *CoordsDB) GetCoordsByAdId(coords *models.Coords) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}

	rows := transaction.QueryRow("SELECT * FROM coords WHERE adid = $1", coords.AdId)
	err = rows.Scan(&coords.Id, &coords.AdId,&coords.X, &coords.Y)
	if err != nil {
		return nil
	}
	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return err
	}
	return nil
}

func (coordsDB *CoordsDB) GetCoords() ([]models.Coords, error) {
	cA := make([]models.Coords, 0)
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return cA, err
	}

	queryString := "SELECT * FROM coords "
	rows, err := transaction.Query(queryString)
	if err != nil {
		return cA, nil
	}
	for rows.Next() {
		var cFound models.Coords
		err = rows.Scan(&cFound.Id, &cFound.AdId, &cFound.X, &cFound.Y)
		if err != nil {
			utils.WriteError(false, "Error scanning row", err)
			errRollback := transaction.Rollback()
			if errRollback != nil {
				utils.WriteError(true, "Error rollback", errRollback)
			}
			return cA, err
		}
		cA = append(cA, cFound)
	}
	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return cA, err
	}
	return cA, nil
}
