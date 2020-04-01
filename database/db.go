package database

import (
	"github.com/jackc/pgx"
)

type Database struct {
	pool *pgx.ConnPool
}

var db Database

func Init(config pgx.ConnConfig) error {
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
	return nil
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

func GetPool() *pgx.ConnPool {
	return db.pool
}