package config

import (
	"fmt"
	"strconv"
)

func Validate() error {

	if !Current.DisableAuth { // If auth is not disabled, validate the master token
		// Validate the master token is set when auth is not disabled
		if Current.MasterToken == "" {
			return fmt.Errorf("master token must be set.")
		}

		// Validate the master token is at least 32 characters long
		if len(Current.MasterToken) < 32 {
			return fmt.Errorf("master token must be at least 32 characters long.")
		}
	}

	// Validate the port is a valid number
	if _, err := strconv.Atoi(Current.Port); err != nil {
		return fmt.Errorf("invalid port number: %s", Current.Port)
	}
	return nil
}
