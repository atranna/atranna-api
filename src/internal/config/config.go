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
}

var Current Config = defaultConfig
