package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DEBUG = true

func NewSQLite() (*sql.DB, error) {
	if DEBUG {
		log.Println("Initilizating DB: ")
	}
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()
	if DEBUG {
		log.Println("DB initilized")
		return nil, err
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db, nil
}

// I have no idea how to run migration till date
// func RunMigrations(db *sql.DB) error {}
