package config

import (
	"fmt"
	"strconv"
)

func Validate() error {

	validateAuthErr := validateAuth()
	if validateAuthErr != nil {
		return validateAuthErr
	}

	validateServiceErr := validateService()
	if validateServiceErr != nil {
		return validateServiceErr
	}

	validateStorageErr := validateStorage()
	if validateStorageErr != nil {
		return validateStorageErr
	}

	return nil
}

func validateAuth() error {
	if Current.Auth.Enable { // If auth is enabled, validate the master token
		if Current.Auth.MasterToken == "" {
			return fmt.Errorf("master token must be set.")
		}

		// Validate the master token is at least 32 characters long
		if len(Current.Auth.MasterToken) < 32 {
			return fmt.Errorf("master token must be at least 32 characters long.")
		}
	}
	return nil
}

func validateService() error {
	// Validate the port is a valid number
	if _, err := strconv.Atoi(Current.Service.Port); err != nil {
		return fmt.Errorf("invalid port number: %s", Current.Service.Port)
	}
	return nil
}

func validateStorage() error {
	// Validate the storage backend is one of the supported backends
	supportedBackends := []string{"memory"}
	isValidBackend := false
	for _, backend := range supportedBackends {
		if Current.Storage.Backend == backend {
			isValidBackend = true
			break
		}
	}
	if !isValidBackend {
		return fmt.Errorf("unsupported storage backend: %s", Current.Storage.Backend)
	}
	return nil
}