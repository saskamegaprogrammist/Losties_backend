package database

import (
	"fmt"
	"github.com/jackc/pgx"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
	"strings"
)

type AdsDB struct {
}

func (adsDB *AdsDB) InsertAd(ad *models.Ad) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}

	row := transaction.QueryRow("INSERT INTO ads (userid, type, title, text, time, contacts, date) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		ad.UserId, ad.Type, ad.Title, ad.Text, ad.Time, ad.Contacts, ad.Date)
	err = row.Scan(&ad.Id)
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

func (adsDB *AdsDB) GetAdById(ad *models.Ad) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}

	//var userPhone sql.NullString
	rows := transaction.QueryRow("SELECT * FROM ads WHERE id = $1", ad.Id)
	err = rows.Scan(&ad.Id, &ad.UserId, &ad.Type, &ad.Title, &ad.Text, &ad.Time, &ad.Contacts, &ad.Comments, &ad.Date)
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

func (adsDB *AdsDB) GetAdsByUserId(userId int, adType int) ([]models.Ad, error) {
	ads := make([]models.Ad, 0)
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return ads, err
	}

	//var userPhone sql.NullString
	rows, err := transaction.Query("SELECT * FROM ads WHERE userid = $1 and type = $2", userId, adType)
	if err != nil {
		utils.WriteError(false, "No ads with this user or wrong type", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			utils.WriteError(true, "Error rollback", errRollback)
		}
		return ads, err
	}
	for rows.Next() {
		var adFound models.Ad
		err = rows.Scan(&adFound.Id, &adFound.UserId, &adFound.Type, &adFound.Title, &adFound.Text, &adFound.Time, &adFound.Contacts, &adFound.Comments, &adFound.Date)
		if err != nil {
			utils.WriteError(false, "Error scanning row", err)
			errRollback := transaction.Rollback()
			if errRollback != nil {
				utils.WriteError(true, "Error rollback", errRollback)
			}
			return ads, err
		}
		ads = append(ads, adFound)
	}
	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return ads, err
	}
	return ads, nil
}
func (adsDB *AdsDB) GetAdsNumberByUserId(userId int, adType int) (int , error) {
	ads := 0
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return ads, err
	}

	var selectString string
	var rows *pgx.Row
	if adType != 0 && adType != 1 {
		selectString = "SELECT COUNT(*) FROM ads WHERE userid = $1 "
		rows = transaction.QueryRow(selectString, userId)
	} else {
		selectString = "SELECT COUNT(*) FROM ads WHERE userid = $1 and type = $2"
		rows = transaction.QueryRow(selectString, userId, adType)
	}
	err = rows.Scan(&ads)
	if err != nil {
		utils.WriteError(false, "Error scanning row", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			utils.WriteError(true, "Error rollback", errRollback)
		}
		return ads, err
	}
	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return ads, err
	}
	return ads, nil
}

func (adsDB *AdsDB) GetAds(adType int) ([]models.Ad, error) {
	return adsDB.getAds(adType, "")
}

func (adsDB *AdsDB) GetAdsSorted(adType int, sort string) ([]models.Ad, error) {
	return adsDB.getAds(adType, sort)
}

func (adsDB *AdsDB) getAds(adType int, sort string) ([]models.Ad, error) {
	ads := make([]models.Ad, 0)
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return ads, err
	}

	var queryString string
	if sort == "" {
		queryString = "SELECT * FROM ads WHERE type = $1"
	} else if sort == "date" {
		queryString = "SELECT * FROM ads WHERE type = $1 ORDER BY date DESC"

	} else if sort == "comments" {
		queryString = "SELECT * FROM ads WHERE type = $1 ORDER BY comments DESC"
	} else {
		errorSort := fmt.Errorf("wrong sort parameter %s", sort)
		utils.WriteError(false, "Wrong sort parameter", errorSort)
		return ads, errorSort
	}

	rows, err := transaction.Query(queryString, adType)
	if err != nil {
		utils.WriteError(false, "Wrong ad type", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			utils.WriteError(true, "Error rollback", errRollback)
		}
		return ads, err
	}
	for rows.Next() {
		var adFound models.Ad
		err = rows.Scan(&adFound.Id, &adFound.UserId, &adFound.Type, &adFound.Title, &adFound.Text, &adFound.Time, &adFound.Contacts, &adFound.Comments, &adFound.Date)
		if err != nil {
			utils.WriteError(false, "Error scanning row", err)
			errRollback := transaction.Rollback()
			if errRollback != nil {
				utils.WriteError(true, "Error rollback", errRollback)
			}
			return ads, err
		}
		ads = append(ads, adFound)
	}
	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return ads, err
	}
	return ads, nil
}

func (adsDB *AdsDB) SearchAds(search string) ([]models.Ad, error) {
	ads := make([]models.Ad, 0)
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return ads, err
	}
	queryString := "SELECT * FROM ads WHERE lower(text) LIKE '%%' || $1 || '%%' OR lower(title) LIKE '%%' || $1 || '%%' OR lower(contacts) LIKE '%%' || $1 || '%%' OR lower(time) LIKE '%%' || $1 || '%%' UNION SELECT * FROM ADS WHERE id IN (SELECT pets.adid FROM pets WHERE lower(pets.name) LIKE '%%' || $1 || '%%' OR lower(pets.animal) LIKE '%%' || $1 || '%%' OR lower(pets.breed) LIKE '%%' || $1 || '%%' OR lower(pets.color) LIKE '%%' || $1 || '%%');"

	rows, err := transaction.Query(queryString, strings.ToLower(search))
	if err != nil {
		utils.WriteError(false, "Wrong ad type", err)
		errRollback := transaction.Rollback()
		if errRollback != nil {
			utils.WriteError(true, "Error rollback", errRollback)
		}
		return ads, err
	}
	for rows.Next() {
		var adFound models.Ad
		err = rows.Scan(&adFound.Id, &adFound.UserId, &adFound.Type, &adFound.Title, &adFound.Text, &adFound.Time, &adFound.Contacts, &adFound.Comments, &adFound.Date)
		if err != nil {
			utils.WriteError(false, "Error scanning row", err)
			errRollback := transaction.Rollback()
			if errRollback != nil {
				utils.WriteError(true, "Error rollback", errRollback)
			}
			return ads, err
		}
		ads = append(ads, adFound)
	}
	err = transaction.Commit()
	if err != nil {
		utils.WriteError(true, "Error commit", err)
		return ads, err
	}
	return ads, nil
}