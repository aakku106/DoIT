package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

