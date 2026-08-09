package devices

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atranna/atranna-api/src/internal/models"
	"github.com/atranna/atranna-api/src/internal/repository/memory"
	"github.com/gin-gonic/gin"
)

func TestGetDevicesUsesOrgHeaderWhenContextOrgMissing(t *testing.T) {
	deviceRepo := memory.NewDeviceRepository()
	_, err := deviceRepo.Add(models.Device{
		Hostname: "PVE-Prod-01",
		IP:       "10.77.0.11",
		Vendor:   "HPE",
		Model:    "Proliant DL360 Gen9",
		Type:     "Server",
		OrgID:    3,
	})
	if err != nil {
		t.Fatalf("failed to seed device repo: %v", err)
	}

	handler := NewHandler(deviceRepo, memory.NewInterfaceRepository(deviceRepo))

	router := gin.Default()
	router.GET("/devices", handler.GetDevices)

	req := httptest.NewRequest(http.MethodGet, "/devices", nil)
	req.Header.Set("X-Org-ID", "3")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var devices []models.Device
	if err := json.Unmarshal(w.Body.Bytes(), &devices); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(devices) != 1 {
		t.Fatalf("expected 1 device, got %d", len(devices))
	}
	if devices[0].OrgID != 3 {
		t.Fatalf("expected org_id 3, got %d", devices[0].OrgID)
	}
}

func TestGetDevicesReturnsEmptyArrayNotNull(t *testing.T) {
	deviceRepo := memory.NewDeviceRepository()
	handler := NewHandler(deviceRepo, memory.NewInterfaceRepository(deviceRepo))

	router := gin.Default()
	router.GET("/devices", handler.GetDevices)

	req := httptest.NewRequest(http.MethodGet, "/devices", nil)
	req.Header.Set("X-Org-ID", "3")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if w.Body.String() != "[]" {
		t.Fatalf("expected empty array response, got %s", w.Body.String())
	}
}
