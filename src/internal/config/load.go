package config

func Load() {
	// Initialize the configuration with default values
	Current = defaultConfig

	// TODO: Load configuration from YAML here

	// Override configuration with environment variables

	Current.Debug = getEnvAsBool("DEBUG", defaultConfig.Debug)
	Current.Auth.Enable = getEnvAsBool("AUTH_ENABLE", defaultConfig.Auth.Enable)
	Current.Auth.MasterToken = getEnv("MASTER_TOKEN", defaultConfig.Auth.MasterToken)
	Current.Service.Port = getEnv("PORT", defaultConfig.Service.Port)
}