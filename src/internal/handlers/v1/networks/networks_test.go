package networks

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

func TestNetworks(t *testing.T) {
	store.Networks = []models.Network{}

	router := gin.Default()
	router.POST("/networks", AddNetwork)
	router.GET("/networks", GetNetworks)
	router.GET("/networks/:id", GetNetwork)
	router.DELETE("/networks/:id", DeleteNetwork)

	// Add a network
	body := map[string]any{
		"name":    "Test-Network-01",
		"cidr":    "192.168.1.0/24",
		"gateway": "192.168.1.1",
		"vlan":    1,
	}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/networks", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	// Get networks
	req = httptest.NewRequest(http.MethodGet, "/networks", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Get network by ID
	req = httptest.NewRequest(http.MethodGet, "/networks/1", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Delete network by ID
	req = httptest.NewRequest(http.MethodDelete, "/networks/1", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", w.Code)
	}
}
