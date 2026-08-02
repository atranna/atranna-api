package database

import (
	"errors"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
)


func ApplySQLiteMigrations() {
	driver, err := sqlite.WithInstance(DB, &sqlite.Config{})
	if err != nil {
		panic(err)
	}	
	m, err := migrate.NewWithDatabaseInstance(
		"file://etc/atranna-api/migrations/sqlite",
		"sqlite3", driver)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}