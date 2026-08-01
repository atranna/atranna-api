package config

import (
	"fmt"
	"os"
	"strconv"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

type Config struct {
	Debug       bool
	DisableAuth bool
	MasterToken string
	Port		string
}

var defaultConfig = Config{
	Debug:       false,
	DisableAuth: false,
	MasterToken: "",
	Port:		"8080",
}

var Current Config = defaultConfig

func Load() {
	// Initialize the configuration with default values
	Current = defaultConfig

	// Load configuration from YAML here

	// Override configuration with environment variables

	Current.Debug = getEnvAsBool("DEBUG", defaultConfig.Debug)
	Current.DisableAuth = getEnvAsBool("DEV_DISABLE_AUTH", defaultConfig.DisableAuth)
	Current.MasterToken = getEnv("MASTER_TOKEN", defaultConfig.MasterToken)
	Current.Port = getEnv("PORT", defaultConfig.Port)
}

func Validate() error {
	// Validate the master token is set when auth is not disabled
	if Current.MasterToken == "" && !Current.DisableAuth {
		return fmt.Errorf("master token must be set.")
	}

	// Validate the port is a valid number
	if _, err := strconv.Atoi(Current.Port); err != nil {
		return fmt.Errorf("invalid port number: %s", Current.Port)
	}
	return nil
}
