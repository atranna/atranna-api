package helpers

import (
	"fmt"
	"time"

	"atranna-api/src/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID int) (string, error) {
	secret := config.Current.Auth.JWTSecret
	if secret == "" {
		return "", fmt.Errorf("JWT secret is not configured")
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"sub": fmt.Sprintf("%d", userID),
		"uid": userID,
		"iat": now.Unix(),
		"exp": now.Add(time.Duration(config.Current.Auth.ExpirationHours) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}