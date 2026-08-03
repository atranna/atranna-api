package helpers

import (
	"atranna-api/src/internal/config"
	"testing"
)

func TestAuthorizationPass(t *testing.T) {
	config := config.Config{
		Auth: struct {
			Enable      bool
			MasterToken string
			JWTSecret   string
			ExpirationHours int
		}{
			Enable:      true,
			MasterToken: "testtoken",
		},
	}

	if !CheckAuthorization("testtoken", config) {
		t.Errorf("expected true, got false")
	}
}

func TestAuthorizationFail(t *testing.T) {
	config := config.Config{
		Auth: struct {
			Enable      bool
			MasterToken string
			JWTSecret   string
			ExpirationHours int
		}{
			Enable:      true,
			MasterToken: "testtoken",
		},
	}

	if CheckAuthorization("wrongtoken", config) {
		t.Errorf("expected false, got true")
	}
}

func TestAuthorizationDisabled(t *testing.T) {
	config := config.Config{
		Auth: struct {
			Enable      bool
			MasterToken string
			JWTSecret   string
			ExpirationHours int
		}{
			Enable: false,
		},
	}

	if !CheckAuthorization("token", config) {
		t.Errorf("expected true, got false")
	}
}

func TestAuthorizationEmptyToken(t *testing.T) {
	config := config.Config{
		Auth: struct {
			Enable      bool
			MasterToken string
			JWTSecret   string
			ExpirationHours int
		}{
			Enable:      true,
			MasterToken: "",
		},
	}

	if CheckAuthorization("token", config) {
		t.Errorf("expected false, got true")
	}

	if CheckAuthorization("", config) {
		t.Errorf("expected false, got true")
	}
}