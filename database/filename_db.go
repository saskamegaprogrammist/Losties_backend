package database

import (
	"context"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

type FilenameDB struct {
}

func (filenameDB *FilenameDB) InsertFilenameAd (filenameAd *models.FilenameAd) error {
	filenameAdCollection := getMongo().Collection("filename_ad")
	_, err := filenameAdCollection.InsertOne(context.TODO(), filenameAd)
	return err
}

func (filenameDB *FilenameDB) GetFilenameByAdId(adId int) (models.FilenameAd, error) {
	cookiesCollection := getMongo().Collection("filename_ad")
	var result models.FilenameAd
	filter := bson.D{{"ad", adId}}
	err := cookiesCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (filenameDB *FilenameDB) InsertFilenameUser (filenameUser *models.FilenameUser) error {
	filenameAdCollection := getMongo().Collection("filename_user")
	_, err := filenameAdCollection.InsertOne(context.TODO(), filenameUser)
	return err
}

func (filenameDB *FilenameDB) GetFilenameByUserId(adId int) (models.FilenameUser, error) {
	cookiesCollection := getMongo().Collection("filename_user")
	var result models.FilenameUser
	filter := bson.D{{"user", adId}}
	err := cookiesCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

