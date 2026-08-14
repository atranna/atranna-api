package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atranna/atranna-api/src/internal/auth"
	"github.com/atranna/atranna-api/src/internal/models"
	"github.com/atranna/atranna-api/src/internal/repository/memory"
	"github.com/gin-gonic/gin"
)

func TestRequireOrganizationPermissionMiddleware(t *testing.T) {
	orgMemberRepo := memory.NewOrganizationMemberRepository()
	orgMemberRepo.Add(models.OrganizationMember{OrganizationID: 1, UserID: 1, Role: "owner"})
	orgMemberRepo.Add(models.OrganizationMember{OrganizationID: 1, UserID: 2, Role: "admin"})
	orgMemberRepo.Add(models.OrganizationMember{OrganizationID: 1, UserID: 3, Role: "viewer"})

	testCases := []struct {
		name       string
		userID     int
		wantStatus int
	}{
		{name: "owner can delete organization", userID: 1, wantStatus: http.StatusOK},
		{name: "admin can delete organization", userID: 2, wantStatus: http.StatusOK},
		{name: "viewer cannot delete organization", userID: 3, wantStatus: http.StatusForbidden},
		{name: "non-member cannot delete organization", userID: 99, wantStatus: http.StatusForbidden},
		{name: "master token cannot delete organization", userID: -1, wantStatus: http.StatusForbidden},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			router := gin.Default()
			router.Use(func(c *gin.Context) {
				c.Set("user_id", tc.userID)
				c.Next()
			})
			router.DELETE("/organizations/:id", RequireOrganizationPermissionMiddleware(orgMemberRepo, auth.OrganizationWrite), func(c *gin.Context) {
				c.Status(http.StatusOK)
			})

			req := httptest.NewRequest(http.MethodDelete, "/organizations/1", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tc.wantStatus {
				t.Fatalf("expected %d, got %d (body: %s)", tc.wantStatus, w.Code, w.Body.String())
			}
		})
	}
}

func TestRequireOrganizationPermissionMiddlewareDifferentOrg(t *testing.T) {
	orgMemberRepo := memory.NewOrganizationMemberRepository()
	orgMemberRepo.Add(models.OrganizationMember{OrganizationID: 1, UserID: 2, Role: "admin"})

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("user_id", 2)
		c.Next()
	})
	router.DELETE("/organizations/:id", RequireOrganizationPermissionMiddleware(orgMemberRepo, auth.OrganizationWrite), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodDelete, "/organizations/2", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Fatalf("expected 403 for a different organization, got %d", w.Code)
	}
}

func TestRequireOrganizationPermissionMiddlewareInvalidOrgID(t *testing.T) {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("user_id", 1)
		c.Next()
	})
	router.DELETE("/organizations/:id", RequireOrganizationPermissionMiddleware(memory.NewOrganizationMemberRepository(), auth.OrganizationWrite), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodDelete, "/organizations/abc", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
