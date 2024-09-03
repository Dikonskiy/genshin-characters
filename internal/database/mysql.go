package database

import (
	"fmt"
	"genshin/internal/config"
	"log"

	"errors"

	"github.com/jmoiron/sqlx"
)

type DB interface {
	GetDB() *sqlx.DB
}

type mysql struct {
	db *sqlx.DB
}

func NewMysqlConn(dbCfg config.Config) (DB, error) {
	log.Println("mysql init...")

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s%s", dbCfg.GetConnection(), "?parseTime=true"))
	if err != nil {
		return nil, errors.Join(errors.New("NewMysqlConn sql.Open"), err)
	}

	return &mysql{
		db: db,
	}, nil
}

func (m mysql) GetDB() *sqlx.DB {
	return m.db
}
