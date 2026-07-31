package interfaces

import (
	devices "atranna-api/src/internal/handlers/v1/devices"
	"atranna-api/src/internal/repository/memory"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestInterfaces(t *testing.T) {
	deviceRepo := memory.NewDeviceRepository()
	interfaceRepo := memory.NewInterfaceRepository(deviceRepo)
	deviceHandler := devices.NewHandler(deviceRepo, interfaceRepo)
	handler := NewHandler(interfaceRepo)

	router := gin.Default()
	router.POST("/devices", deviceHandler.Add)

	router.POST("/interfaces/", handler.Add)
	router.GET("/interfaces", handler.GetInterfaces)
	router.GET("/interfaces/:id", handler.GetInterface)
	router.DELETE("/interfaces/:id", handler.Delete)

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

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
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

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
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

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", w.Code)
	}
}

func TestAddInterfaceRequiresDeviceID(t *testing.T) {
	deviceRepo := memory.NewDeviceRepository()
	handler := NewHandler(memory.NewInterfaceRepository(deviceRepo))

	router := gin.Default()
	router.POST("/interfaces/", handler.Add)

	body := map[string]any{
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
