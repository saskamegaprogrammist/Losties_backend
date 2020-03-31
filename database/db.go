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
	_, err := database.pool.Exec(``)
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