package database

import (
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
)

type CommentsDB struct {
}

func (commentsDB *CommentsDB) InsertComment(comment *models.Comment) error {
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return err
	}

	row := transaction.QueryRow("INSERT INTO comments (adid, userid, text, date) VALUES ($1, $2, $3, $4) RETURNING id",
		comment.AdId, comment.UserId, comment.Text, comment.Date)
	err = row.Scan(&comment.Id)
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

func (commentsDB *CommentsDB) GetCommentsByAdId(adId int) ([]models.Comment, error) {
	cA := make([]models.Comment, 0)
	db := getPool()
	transaction, err := db.Begin()
	if err != nil {
		utils.WriteError(false, "Failed to start transaction", err)
		return cA, err
	}

	queryString := "SELECT * FROM comments WHERE adid = $1 "
	rows, err := transaction.Query(queryString, adId)
	if err != nil {
		return cA, nil
	}
	for rows.Next() {
		var cFound models.Comment
		err = rows.Scan(&cFound.Id, &cFound.AdId, &cFound.UserId, &cFound.Text, &cFound.Date)
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