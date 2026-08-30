package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	schema "github.com/aakku106/DoIT/sql"
	_ "github.com/ncruces/go-sqlite3/driver"
)

// NewSQLite is called by main.go for subcommands (add, list, etc.)
// It automatically finds .doit in the current directory or parent folders.
func NewSQLite() (*sql.DB, error) {
	doitDir, err := findDoitDir()
	if err != nil {
		return nil, err
	}

	dbPath := filepath.Join(doitDir, "doit.db")
	dsn := fmt.Sprintf("file:%s?_journal_mode=WAL&_foreign_keys=on", dbPath)

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// InitSQLite is called strictly by initProject() to build .doit/doit.db & run schema
func InitSQLite(doitDir string) (*sql.DB, error) {
	dbPath := filepath.Join(doitDir, "doit.db")
	dsn := fmt.Sprintf("file:%s?_journal_mode=WAL&_foreign_keys=on", dbPath)

	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	if _, err := db.Exec(schema.SchemaSQLite); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// Helper function to search upwards for the .doit folder
func findDoitDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := cwd
	for {
		doitPath := filepath.Join(dir, ".doit")
		if info, err := os.Stat(doitPath); err == nil && info.IsDir() {
			return doitPath, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir { // Reached system root
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("not a doit repository (run 'doit init' first)")
}
