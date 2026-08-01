package helpers

import (
	"atranna-api/src/internal/config"
	"crypto/subtle"
	"strings"
)

func CheckAuthorization(authorizationHeader string, config config.Config) bool {
	if !config.Auth.Enable {
		return true
	}

	masterToken := strings.TrimSpace(config.Auth.MasterToken)
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
