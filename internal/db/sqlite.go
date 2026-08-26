package db

import (
	"database/sql"
	"fmt"
	"path/filepath"

	schema "github.com/aakku106/DoIT/sql"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

// InitSQLite initializes the db at a specific path and applies schema
func InitSQLite(doitDir string) (*sql.DB, error) {
	dbPath := filepath.Join(doitDir, "doit.db")

	// Format as SQLite DSN URI
	dsn := fmt.Sprintf("file:%s?_journal_mode=WAL&_foreign_keys=on", dbPath)

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	// Apply schema migrations
	if _, err := db.Exec(schema.SchemaSQLite); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
