package helpers

import (
	"fmt"
	"strconv"

	"atranna-api/src/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateJWT(tokenString string) (int, error) {
	secret := config.Current.Auth.JWTSecret
	if secret == "" {
		return 0, fmt.Errorf("JWT secret is not configured")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	if uid, ok := claims["uid"]; ok {
		switch v := uid.(type) {
		case float64:
			return int(v), nil
		case int:
			return v, nil
		case int64:
			return int(v), nil
		case string:
			parsed, convErr := strconv.Atoi(v)
			if convErr != nil {
				return 0, fmt.Errorf("invalid uid claim")
			}
			return parsed, nil
		}
	}

	if sub, ok := claims["sub"].(string); ok {
		parsed, convErr := strconv.Atoi(sub)
		if convErr != nil {
			return 0, fmt.Errorf("invalid sub claim")
		}
		return parsed, nil
	}

	return 0, fmt.Errorf("missing user id claim")
}