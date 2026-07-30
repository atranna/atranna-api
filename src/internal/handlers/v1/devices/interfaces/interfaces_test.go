package v1

import (
	devices "atranna-api/src/internal/handlers/v1/devices"
	"atranna-api/src/internal/models"
	"atranna-api/src/internal/store"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestInterfaces(t *testing.T) {
	store.Devices = []models.Device{}
	store.Interfaces = []models.Interface{}

	router := gin.Default()
	router.POST("/devices", devices.AddDevice)

	router.POST("/interfaces/", AddInterface)
	router.GET("/interfaces", GetInterfaces)
	router.GET("/interfaces/:id", GetInterface)
	router.DELETE("/interfaces/:interface_id", DeleteInterface)

	// Add a device
	body := map[string]any{
		"hostname": "Test-Device-01",
		"ip":       "192.168.1.100",
		"vendor":   "Unknown",
		"model":    "Fake",
		"type":     "Test-Device",
	}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/devices", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Add an interface
	body = map[string]any{
		"device_id":   1,
		"name":        "eth0",
		"ip_address":  "192.168.1.100",
		"mac_address": "00:11:22:33:44:55",
		"state":       "up",
		"speed":       1000,
	}
	bodyJSON, err = json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req = httptest.NewRequest(http.MethodPost, "/interfaces/", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Get interface
	req = httptest.NewRequest(http.MethodGet, "/interfaces/1", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Delete interface by ID
	req = httptest.NewRequest(http.MethodDelete, "/interfaces/1", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
}

func TestAddInterfaceRequiresDeviceID(t *testing.T) {
	t.Setenv("DEV_DISABLE_AUTH", "true")

	store.Devices = []models.Device{}
	store.Interfaces = []models.Interface{}

	router := gin.Default()
	router.POST("/interfaces/", AddInterface)

	body := map[string]any{
		"device_id": 1,
		"name": "eth0",
	}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/interfaces/", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
