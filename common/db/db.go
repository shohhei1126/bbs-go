package db

import (
	"database/sql"
)

func NewMySqlDb(dbString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
