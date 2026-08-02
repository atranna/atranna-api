package config

import (
	"fmt"
	"strconv"
)

func Validate() error {

	validateAuthErr := validateAuth(Current)
	if validateAuthErr != nil {
		return validateAuthErr
	}

	validateServiceErr := validateService(Current)
	if validateServiceErr != nil {
		return validateServiceErr
	}

	validateStorageErr := validateStorage(Current)
	if validateStorageErr != nil {
		return validateStorageErr
	}

	return nil
}

func validateAuth(cfg Config) error {
	if cfg.Auth.Enable { // If auth is enabled, validate the master token
		if cfg.Auth.MasterToken == "" {
			return fmt.Errorf("master token must be set.")
		}

		// Validate the master token is at least 32 characters long
		if len(cfg.Auth.MasterToken) < 32 {
			return fmt.Errorf("master token must be at least 32 characters long.")
		}
	}
	return nil
}

func validateService(cfg Config) error {
	// Validate the port is a valid number
	if _, err := strconv.Atoi(cfg.Service.Port); err != nil {
		return fmt.Errorf("invalid port number: %s", cfg.Service.Port)
	}
	return nil
}

func validateStorage(cfg Config) error {
	// Validate the storage backend is one of the supported backends
	supportedBackends := []string{"memory", "postgres", "sqlite"}
	isValidBackend := false
	for _, backend := range supportedBackends {
		if cfg.Storage.Backend == backend {
			isValidBackend = true
			break
		}
	}
	if !isValidBackend {
		return fmt.Errorf("unsupported storage backend: %s", cfg.Storage.Backend)
	}
	
	if cfg.Storage.Backend == "postgres" {
		// Validate that all required Postgres fields are set
		if cfg.Storage.Postgres.Host == "" {
			return fmt.Errorf("Postgres host must be set.")
		}
		if cfg.Storage.Postgres.Port == "" {
			return fmt.Errorf("Postgres port must be set.")
		}
		if cfg.Storage.Postgres.User == "" {
			return fmt.Errorf("Postgres user must be set.")
		}
		if cfg.Storage.Postgres.Password == "" {
			return fmt.Errorf("Postgres password must be set.")
		}
		if cfg.Storage.Postgres.DBName == "" {
			return fmt.Errorf("Postgres database name must be set.")
		}
	}

	if cfg.Storage.Backend == "sqlite" {
		// Validate that the SQLite file path is set
		if cfg.Storage.SQLite.FilePath == "" {
			return fmt.Errorf("SQLite file path must be set.")
		}
	}

	return nil
}