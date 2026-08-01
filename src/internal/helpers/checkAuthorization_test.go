package helpers

import (
	"atranna-api/src/internal/config"
	"testing"
)

func TestAuthorizationPass(t *testing.T) {
	config := config.Config{
		MasterToken: "testtoken",
	}

	if !CheckAuthorization("testtoken", config) {
		t.Errorf("expected true, got false")
	}
}

func TestAuthorizationFail(t *testing.T) {
	config := config.Config{
		MasterToken: "testtoken",
	}

	if CheckAuthorization("wrongtoken", config) {
		t.Errorf("expected false, got true")
	}
}

func TestAuthorizationDisabled(t *testing.T) {
	config := config.Config{
		DisableAuth: true,
	}

	if !CheckAuthorization("token", config) {
		t.Errorf("expected true, got false")
	}
}

func TestAuthorizationEmptyToken(t *testing.T) {
	config := config.Config{
		MasterToken: "",
	}

	if CheckAuthorization("token", config) {
		t.Errorf("expected false, got true")
	}

	if CheckAuthorization("", config) {
		t.Errorf("expected false, got true")
	}
}