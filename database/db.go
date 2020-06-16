package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"github.com/saskamegaprogrammist/Losties_backend/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	pool *pgx.ConnPool
	UsersDB *UsersDB
	CookiesDB *CookiesDB
	FilenameDB *FilenameDB
	CoommentsDB *CommentsDB
	AdsDB *AdsDB
	PetsDB *PetsDB
	CoordsDB *CoordsDB
	mongoDB *mongo.Database
}

var db Database

func Init(config pgx.ConnConfig, mongoHost string) error {
	InitMongoDb(mongoHost)
	var err error
	db.pool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: config,
		MaxConnections: 20,
	})
	if err != nil {
		return err
	}
	//err = db.createTables()
	if err != nil {
		return err
	}
	db.UsersDB = &UsersDB{}
	db.AdsDB = &AdsDB{}
	db.PetsDB = &PetsDB{}
	db.CoommentsDB = &CommentsDB{}
	db.CoordsDB = &CoordsDB{}
	return nil
}

func InitMongoDb (host string) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:27017", host))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		utils.WriteError(true, "Can't connect to mongodb", err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		utils.WriteError(true, "Can't ping mongodb", err)
	}
	db.mongoDB = client.Database("losties")
	//err = db.mongoDB.Collection("cookies").Drop(context.TODO())
	//if err != nil {
	//	utils.WriteError(true, "Can't drop cookies collection", err)
	//}
	//err = db.mongoDB.Collection("filename_ad").Drop(context.TODO())
	//if err != nil {
	//	utils.WriteError(true, "Can't drop filename_ad collection", err)
	//}
	//err = db.mongoDB.Collection("filename_user").Drop(context.TODO())
	//if err != nil {
	//	utils.WriteError(true, "Can't drop filename_user collection", err)
	//}
}

func (database *Database)  createTables() error {
	_, err := database.pool.Exec(`
DROP TABLE IF EXISTS pets;
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS coords;
DROP TABLE IF EXISTS ads;
DROP TABLE IF EXISTS users;


CREATE TABLE users (
    id SERIAL NOT NULL PRIMARY KEY,
    firstname text NOT NULL ,
    lastname text NOT NULL,
    email text NOT NULL UNIQUE,
    nickname text NOT NULL UNIQUE,
    phone text UNIQUE,
    password text NOT NULL
    CONSTRAINT valid_email CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$')
);


CREATE TABLE ads (
    id SERIAL NOT NULL PRIMARY KEY,
    userid int NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    type int NOT NULL  CHECK (type = 0 OR type = 1),
    title text DEFAULT '',
    text text DEFAULT '',
    time text DEFAULT '',
    contacts text DEFAULT '',
    comments int DEFAULT 0,
    date TIMESTAMPTZ NOT NULL
);

CREATE TABLE pets (
    id SERIAL NOT NULL PRIMARY KEY,
    adid int NOT NULL REFERENCES ads(id) ON DELETE CASCADE,
    name text DEFAULT '',
    animal text DEFAULT '',
    breed text DEFAULT '',
    color text DEFAULT ''
);

CREATE TABLE coords (
    id SERIAL NOT NULL PRIMARY KEY,
    adid int NOT NULL REFERENCES ads(id) ON DELETE CASCADE,
    x double precision NOT NULL ,
    y double precision NOT NULL 
);

CREATE TABLE comments (
    id SERIAL NOT NULL PRIMARY KEY,
    adid int NOT NULL REFERENCES ads(id) ON DELETE CASCADE,
    userid int NOT NULL REFERENCES users(id) ON DELETE SET NULL,
    text text NOT NULL ,
	date TIMESTAMPTZ NOT NULL
);


CREATE OR REPLACE FUNCTION ad_comment() RETURNS TRIGGER
LANGUAGE  plpgsql
AS $ad_comment$
BEGIN
   UPDATE ads SET comments = comments + 1 WHERE id = NEW.adid;
    RETURN NEW;
END
$ad_comment$;


CREATE TRIGGER AdComment
    AFTER INSERT on comments
    FOR EACH ROW
    EXECUTE PROCEDURE ad_comment();


CREATE OR REPLACE FUNCTION delete_comment() RETURNS TRIGGER
LANGUAGE  plpgsql
AS $ad_comment$
BEGIN
   UPDATE ads SET comments = comments - 1 WHERE id = OLD.adid;
    RETURN NEW;
END
$ad_comment$;


CREATE TRIGGER DeleteComment
    AFTER DELETE on comments
    FOR EACH ROW
    EXECUTE PROCEDURE delete_comment();

`)
	if err != nil {
		return err
	}
	return nil
}

func (database *Database)  Destroy() error {
	_, err := database.pool.Exec(``)
	if err != nil {
		return err
	}
	return nil
}

func getPool() *pgx.ConnPool {
	return db.pool
}

func getMongo() *mongo.Database {
	return db.mongoDB
}

func GetUsersDB() *UsersDB {
	return db.UsersDB
}

func GetCookiesDB() *CookiesDB {
	return db.CookiesDB
}

func GetCommentsDB() *CommentsDB {
	return db.CoommentsDB
}


func GetFilenameDB() *FilenameDB {
	return db.FilenameDB
}

func GetAdsDB() *AdsDB {
	return db.AdsDB
}

func GetPetsDB() *PetsDB {
	return db.PetsDB
}

func GetCoordsDB() *CoordsDB {
	return db.CoordsDB
}