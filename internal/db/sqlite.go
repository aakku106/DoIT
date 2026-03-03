package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func NewSQLite() {
	log.Println("Func called")
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
}
