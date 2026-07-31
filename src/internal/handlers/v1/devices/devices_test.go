package devices

import (
	"atranna-api/src/internal/models"
	"atranna-api/src/internal/store"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDevices(t *testing.T) {
	store.Devices = []models.Device{}

	router := gin.Default()
	router.POST("/devices", AddDevice)
	router.GET("/devices", GetDevices)
	router.GET("/devices/:id", GetDevice)
	router.DELETE("/devices/:id", DeleteDevice)

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

	// Get devices
	req = httptest.NewRequest(http.MethodGet, "/devices", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Get device by ID
	req = httptest.NewRequest(http.MethodGet, "/devices/1", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Delete device by ID
	req = httptest.NewRequest(http.MethodDelete, "/devices/1", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", w.Code)
	}

	// Verify device is deleted
	req = httptest.NewRequest(http.MethodGet, "/devices/1", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}
