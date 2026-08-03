package main

import (
	"github.com/atranna/atranna-api/src/internal/config"
	"github.com/atranna/atranna-api/src/internal/router"
)


func main() {
	config.Load()
	err := config.Validate();
	if err != nil {
		panic(err)
	}
	
	router := router.New()
	router.Run("0.0.0.0:" + config.Current.Service.Port)
}
