package organizations

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atranna/atranna-api/src/internal/repository/memory"

	"github.com/gin-gonic/gin"
)

func TestOrganizations(t *testing.T) {
	organizationRepo := memory.NewOrganizationRepository()
	organizationMemberRepo := memory.NewOrganizationMemberRepository()
	handler := NewHandler(organizationRepo, organizationMemberRepo)

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("user_id", 1)
		c.Next()
	})
	router.POST("/organizations", handler.Add)
	router.GET("/organizations", handler.GetOrganizations)
	router.GET("/organizations/:id", handler.GetOrganization)
	router.DELETE("/organizations/:id", handler.Delete)

	// Add a organizations
	body := map[string]any{
		"name": "Test Organization",
		"slug": "test-organization",
	}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/organizations", bytes.NewReader(bodyJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	// Get organizations
	req = httptest.NewRequest(http.MethodGet, "/organizations", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Get organizations by ID
	req = httptest.NewRequest(http.MethodGet, "/organizations/1", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// Delete organizations by ID
	req = httptest.NewRequest(http.MethodDelete, "/organizations/1", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204, got %d", w.Code)
	}

	// Verify organization is deleted
	req = httptest.NewRequest(http.MethodGet, "/organizations/1", nil)
	w = httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", w.Code)
	}
}
