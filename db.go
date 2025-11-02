package main

import (
	"database/sql"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

func openDB() (*sql.DB, error) {
	_ = os.MkdirAll("data", 0755)
	db, err := sql.Open("sqlite", "data/app.db")
	if err != nil {
		return nil, err
	}
	// Apply schema
	if err := applySQL(db, "sql/schema.sql"); err != nil {
		return nil, err
	}
	// Seed only if empty
	var count int
	_ = db.QueryRow("SELECT COUNT(*) FROM todos").Scan(&count)
	if count == 0 {
		_ = applySQL(db, "sql/seed.sql")
	}
	return db, nil
}

func applySQL(db *sql.DB, path string) error {
	b, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return err
	}
	_, err = db.Exec(string(b))
	return err
}
