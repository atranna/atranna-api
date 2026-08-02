package database

import (
	"atranna-api/src/internal/config"
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

func ConnectSQLite() (*sql.DB, error) {
	dsn := config.Current.Storage.SQLite.FilePath + "?_pragma=foreign_keys(1)"

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to open sqlite database: %w", err)
	}
	return db, nil
}