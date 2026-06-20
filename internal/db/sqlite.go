package db

import (
	"database/sql"

	schema "github.com/aakku106/DoIT/sql"
	_ "github.com/ncruces/go-sqlite3/driver"
)

func NewSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if _, err := db.Exec(schema.SchemaSQLite); err != nil {
		return nil, err
	}

	return db, nil
}
