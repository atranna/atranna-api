package database

import (
	"atranna-api/src/internal/config"
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitPostgres() {
	if config.Current.Storage.Backend == "postgres" {
		var err error
		DB, err = ConnectPostgres()
		if err != nil {
			panic(err)
		}
	} else {
		panic("PostgreSQL backend is not enabled in the configuration")
	}
}


func ConnectPostgres() (*sql.DB, error) {

	host := config.Current.Storage.Postgres.Host
	port := config.Current.Storage.Postgres.Port
	user := config.Current.Storage.Postgres.User
	password := config.Current.Storage.Postgres.Password
	dbname := config.Current.Storage.Postgres.DBName

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	return sql.Open("postgres", dsn)
}

func GetPostgresVersion() (string, error) {
	var version string
	err := DB.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return "", err
	}
	return version, nil
}

func ApplyMigrations() {
	driver, err := postgres.WithInstance(DB, &postgres.Config{})
	if err != nil {
		panic(err)
	}	
	m, err := migrate.NewWithDatabaseInstance(
        "file://etc/atranna-api/migrations",
        "postgres", driver)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		panic(err)
	}
}