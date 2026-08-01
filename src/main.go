package main

import (
	"atranna-api/src/internal/config"
	"atranna-api/src/internal/router"
)


func main() {
	config.Load()
	router := router.New()
	router.Run("0.0.0.0:" + config.Current.Port)
}
