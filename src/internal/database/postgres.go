package database

import (
	"database/sql"
	"fmt"

	"github.com/atranna/atranna-api/src/internal/config"

	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)


func ConnectPostgres() (*sql.DB, error) {

	host := config.Current.Storage.Postgres.Host
	port := config.Current.Storage.Postgres.Port
	user := config.Current.Storage.Postgres.User
	password := config.Current.Storage.Postgres.Password
	dbname := config.Current.Storage.Postgres.DBName

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to open postgres database: %w", err)
	}
	return db, nil
}

func GetPostgresVersion() (string, error) {
	var version string
	err := DB.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		return "", err
	}
	return version, nil
}