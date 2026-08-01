package config

func Load() {
	// Initialize the configuration with default values
	Current = defaultConfig

	// TODO: Load configuration from YAML here

	// Override configuration with environment variables

	Current.Debug = getEnvAsBool("DEBUG", defaultConfig.Debug)
	Current.Auth.Enable = getEnvAsBool("AUTH_ENABLE", defaultConfig.Auth.Enable)
	Current.Auth.MasterToken = getEnv("AUTH_MASTER_TOKEN", defaultConfig.Auth.MasterToken)
	Current.Service.Port = getEnv("SERVICE_PORT", defaultConfig.Service.Port)
	Current.Storage.Backend = getEnv("STORAGE_BACKEND", defaultConfig.Storage.Backend)
	Current.Storage.Postgres.Host = getEnv("STORAGE_POSTGRES_HOST", defaultConfig.Storage.Postgres.Host)
	Current.Storage.Postgres.Port = getEnv("STORAGE_POSTGRES_PORT", defaultConfig.Storage.Postgres.Port)
	Current.Storage.Postgres.User = getEnv("STORAGE_POSTGRES_USER", defaultConfig.Storage.Postgres.User)
	Current.Storage.Postgres.Password = getEnv("STORAGE_POSTGRES_PASSWORD", defaultConfig.Storage.Postgres.Password)
	Current.Storage.Postgres.DBName = getEnv("STORAGE_POSTGRES_DBNAME", defaultConfig.Storage.Postgres.DBName)
}