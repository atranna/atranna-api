package config

var defaultConfig = Config{
	Debug:       false,
	Service: struct {
		Port string
	}{
		Port: "8080",
	},
	CORS: struct {
		AllowedOrigins string
	}{
		AllowedOrigins: "http://localhost:5173,http://127.0.0.1:5173,http://localhost:4173,http://127.0.0.1:4173",
	},
	Auth: struct {
		Enable      bool
		MasterToken string
		JWTSecret   string
		ExpirationHours int
		UserCreationEnabled bool
	}{
		Enable:      true,
		MasterToken: "",
		JWTSecret:   "",
		ExpirationHours: 24,
		UserCreationEnabled: true,
	},
	Storage: struct {
		Backend string
		Postgres struct {
			Host     string
			Port     string
			User     string
			Password string
			DBName   string
		}
		SQLite struct {
			FilePath string
		}
	}{
		Backend: "sqlite",
		Postgres: struct {
			Host     string
			Port     string
			User     string
			Password string
			DBName   string
		}{
			Host:     "",
			Port:     "5432",
			User:     "atranna",
			Password: "",
			DBName:   "atranna-api",
		},
		SQLite: struct {
			FilePath string
		}{
			FilePath: "/etc/atranna-api/data/atranna-api.db",
		},
	},
}