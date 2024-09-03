package database

import (
	"database/sql"
	"genshin/internal/config"
	"log"

	"errors"
)

type DB interface {
	GetDB() *sql.DB
}

type mysql struct {
	db *sql.DB
}

func NewMysqlConn(dbCfg config.Config) (DB, error) {
	log.Println("mysql init...")

	db, err := sql.Open("mysql", dbCfg.GetConnection())
	if err != nil {
		return nil, errors.Join(errors.New("NewMysqlConn sql.Open"), err)
	}

	return &mysql{
		db: db,
	}, nil
}

func (p mysql) GetDB() *sql.DB {
	return p.db
}
