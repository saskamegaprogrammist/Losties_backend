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
	mongoDB *mongo.Database
}

var db Database

func Init(config pgx.ConnConfig, mongoHost string) error {
	InitCookiesDb(mongoHost)
	var err error
	db.pool, err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: config,
		MaxConnections: 20,
	})
	if err != nil {
		return err
	}
	err = db.createTables()
	if err != nil {
		return err
	}
	db.UsersDB = &UsersDB{}
	return nil
}

func InitCookiesDb (host string) {
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
	err = db.mongoDB.Collection("cookies").Drop(context.TODO())
	if err != nil {
		utils.WriteError(true, "Can't drop cookies collection", err)
	}
}

func (database *Database)  createTables() error {
	_, err := database.pool.Exec(`
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