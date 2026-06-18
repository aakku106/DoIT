package db

import (
	"database/sql"
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

	schema := `
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY,
		session TEXT NOT NULL DEFAULT 'todo',
		title TEXT NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		expires_at DATETIME,
		completed INTEGER NOT NULL DEFAULT 0 CHECK (completed IN (0,1))
	);`

	if _, err := db.Exec(schema); err != nil {
		return nil, err
	}

	return db, nil
}
