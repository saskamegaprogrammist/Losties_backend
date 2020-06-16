package database

import (
	"context"
	"github.com/saskamegaprogrammist/Losties_backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

type CookiesDB struct {
}

func (cookiesDB *CookiesDB) InsertCookie (cookie *models.CookieInfo) error {
	//log.Println(cookie)
	cookiesCollection := getMongo().Collection("cookies")
	_, err := cookiesCollection.InsertOne(context.TODO(), cookie)
	return err
}

func (cookiesDB *CookiesDB) DeleteCookie (cookie string) error {
	cookiesCollection := getMongo().Collection("cookies")
	filter := bson.D{{"cookie", cookie}}
	_, err := cookiesCollection.DeleteOne(context.TODO(), filter)
	return err
}

func (cookiesDB *CookiesDB) GetCookieByUserId (userId string) (models.CookieInfo, error) {
	cookiesCollection := getMongo().Collection("cookies")
	var result models.CookieInfo
	filter := bson.D{{"user", userId}}
	err := cookiesCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (cookiesDB *CookiesDB) GetUserIdByCookie (cookie string) (models.CookieInfo, error) {
	//log.Println(cookie)
	cookiesCollection := getMongo().Collection("cookies")
	var result models.CookieInfo
	filter := bson.D{{"cookie", cookie}}
	err := cookiesCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}