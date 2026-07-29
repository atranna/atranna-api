package helpers

import (
	"os"
	"strings"
)

func CheckAuthorization(token string) bool {
	if os.Getenv("DEV_DISABLE_AUTH") == "true" {
		return true
	}

	masterToken := strings.TrimSpace(os.Getenv("MASTER_TOKEN"))

	token = strings.TrimSpace(token)
	token = strings.TrimPrefix(token, "Bearer ")

	return token == masterToken
}