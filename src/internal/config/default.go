package config

var defaultConfig = Config{
	Debug:       false,
	Service: struct {
		Port string
	}{
		Port: "8080",
	},
	Auth: struct {
		Enable      bool
		MasterToken string
	}{
		Enable:      true,
		MasterToken: "",
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
		Backend: "memory",
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