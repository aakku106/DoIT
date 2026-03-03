package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DEBUG = true

func NewSQLite(path string) (*sql.DB, error) {
	if DEBUG {
		log.Println("Initilizating DB: ", path)
	}
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if DEBUG {
		log.Println("DB initilized")
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db, err
}
