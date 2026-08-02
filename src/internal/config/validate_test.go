package config

import "testing"

func TestAuthValidationBlankMasterToken(t *testing.T) {
	config := defaultConfig
	
	config.Auth.MasterToken = ""

	err := validateAuth(config)
	if err == nil {
		t.Errorf("Expected error for empty master token, got nil")
	}
}

func TestAuthValidationShortMasterToken(t *testing.T) {
	config := defaultConfig
	
	config.Auth.MasterToken = "short"

	err := validateAuth(config)
	if err == nil {
		t.Errorf("Expected error for short master token, got nil")
	}
}

func TestAuthValidationValidMasterToken(t *testing.T) {
	config := defaultConfig
	
	config.Auth.MasterToken = "thisisavalidmastertokenthatis32chars"

	err := validateAuth(config)
	if err != nil {
		t.Errorf("Expected no error for valid master token, got %v", err)
	}
}

func TestServiceValidationInvalidPort(t *testing.T) {
	config := defaultConfig
	config = Config{
		Service: struct {
			Port string
		}{
			Port: "invalid",
		},
	}
	err := validateService(config)
	if err == nil {
		t.Errorf("Expected error for invalid port, got nil")
	}
}

func TestServiceValidationValidPort(t *testing.T) {
	config := defaultConfig
	config = Config{
		Service: struct {
			Port string
		}{
			Port: "8080",
		},
	}
	err := validateService(config)
	if err != nil {
		t.Errorf("Expected no error for valid port, got %v", err)
	}
}

func TestStorageValidationUnsupportedBackend(t *testing.T) {
	config := defaultConfig

	config.Storage.Backend = "unsupported"

	err := validateStorage(config)
	if err == nil {
		t.Errorf("Expected error for unsupported storage backend, got nil")
	}
}

func TestStorageValidationSupportedBackend(t *testing.T) {
	config := defaultConfig

	config.Storage.Backend = "memory"

	err := validateStorage(config)
	if err != nil {
		t.Errorf("Expected no error for supported storage backend, got %v", err)
	}
}
