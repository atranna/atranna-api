package config

func Load() {
	// Initialize the configuration with default values
	Current = defaultConfig

	// TODO: Load configuration from YAML here

	// Override configuration with environment variables

	Current.Debug = getEnvAsBool("DEBUG", defaultConfig.Debug)
	Current.DisableAuth = getEnvAsBool("DEV_DISABLE_AUTH", defaultConfig.DisableAuth)
	Current.MasterToken = getEnv("MASTER_TOKEN", defaultConfig.MasterToken)
	Current.Port = getEnv("PORT", defaultConfig.Port)
}