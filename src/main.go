package main

import (
	"atranna-api/src/internal/router"
)

func main() {
	router := router.New()
	router.Run("0.0.0.0:8080")
}
