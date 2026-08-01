package config

type Config struct {
	Debug       bool
	Service struct {
		Port string
	}
	Auth struct {
		Enable bool
		MasterToken string
	}
	Storage struct {
		Backend string
		Postgres struct {
			Host     string
			Port     string
			User     string
			Password string
			DBName   string
		}
	}
}

var Current Config = defaultConfig
