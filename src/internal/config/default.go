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
}