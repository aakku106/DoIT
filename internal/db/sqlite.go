package db

import (
	"database/sql"

	schema "github.com/aakku106/DoIT/sql"
	_ "github.com/ncruces/go-sqlite3/driver"
)

const DataBaseName string = "doit.db"

func NewSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DataBaseName) // db name is changed from todo.db to doit.db
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
