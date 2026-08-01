package config

type Config struct {
	Debug       bool
	DisableAuth bool
	MasterToken string
	Port		string
}

var Current Config = defaultConfig
