package main

import (
	"atranna-api/src/internal/config"
	"atranna-api/src/internal/router"
)


func main() {
	config.Load()
	err := config.Validate();
	if err != nil {
		panic(err)
	}
	
	router := router.New()
	router.Run("0.0.0.0:" + config.Current.Port)
}
