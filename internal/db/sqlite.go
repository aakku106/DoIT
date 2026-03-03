package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aakku106/DoIT/internal/store"
	_ "github.com/mattn/go-sqlite3"
)

func NewSQLite() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "todo.db")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()
	queries := store.New(db)
	fmt.Printf("queries: %v\n", queries)
	return db, nil
}
