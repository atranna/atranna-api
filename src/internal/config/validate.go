package config

import (
	"fmt"
	"strconv"
)

func Validate() error {

	if Current.Auth.Enable { // If auth is enabled, validate the master token
		if Current.Auth.MasterToken == "" {
			return fmt.Errorf("master token must be set.")
		}

		// Validate the master token is at least 32 characters long
		if len(Current.Auth.MasterToken) < 32 {
			return fmt.Errorf("master token must be at least 32 characters long.")
		}
	}

	// Validate the port is a valid number
	if _, err := strconv.Atoi(Current.Service.Port); err != nil {
		return fmt.Errorf("invalid port number: %s", Current.Service.Port)
	}
	return nil
}
