package helpers

import (
	"crypto/subtle"
	"os"
	"strings"
)

func CheckAuthorization(authorizationHeader string) bool {
	if os.Getenv("DEV_DISABLE_AUTH") == "true" {
		return true
	}

	masterToken := strings.TrimSpace(os.Getenv("MASTER_TOKEN"))
	if masterToken == "" {
		return false
	}

	token := strings.TrimSpace(authorizationHeader)
	if strings.HasPrefix(strings.ToLower(token), "bearer ") {
		token = token[len("bearer "):]
	}
	token = strings.TrimSpace(token)
	if token == "" {
		return false
	}

	return subtle.ConstantTimeCompare([]byte(token), []byte(masterToken)) == 1
}