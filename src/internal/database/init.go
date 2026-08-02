package database

import (
	"atranna-api/src/internal/config"
	"database/sql"
)

var DB *sql.DB

func Init() {
	switch config.Current.Storage.Backend {
	case "postgres":
		var err error
		DB, err = ConnectPostgres()
		if err != nil {
			panic(err)
		}
	case "sqlite":
		var err error
		DB, err = ConnectSQLite()
		if err != nil {
			panic(err)
		}
	} 
}