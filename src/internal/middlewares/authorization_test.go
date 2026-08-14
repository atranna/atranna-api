package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/atranna/atranna-api/src/internal/models"
	"github.com/atranna/atranna-api/src/internal/repository/memory"
	"github.com/gin-gonic/gin"
)

func setupAuthorizationRouter(orgMemberRepo *memory.OrganizationMemberRepository) *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("user_id", 1)
		c.Next()
	})
	router.GET("/test", AuthorizationMiddleware(orgMemberRepo), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"org_id": c.GetInt("org_id"), "role": c.GetString("role")})
	})
	return router
}

func performAuthorizationRequest(router *gin.Engine, orgHeader string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	if orgHeader != "" {
		req.Header.Set("X-Org-ID", orgHeader)
	}
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestAuthorizationMiddlewareSingleMembershipWithoutHeader(t *testing.T) {
	repo := memory.NewOrganizationMemberRepository()
	repo.Add(models.OrganizationMember{OrganizationID: 1, UserID: 1, Role: "owner"})

	w := performAuthorizationRequest(setupAuthorizationRouter(repo), "")

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if w.Body.String() != `{"org_id":1,"role":"owner"}` {
		t.Fatalf("unexpected response: %s", w.Body.String())
	}
}

func TestAuthorizationMiddlewareHonorsExplicitOrgHeader(t *testing.T) {
	repo := memory.NewOrganizationMemberRepository()
	repo.Add(models.OrganizationMember{OrganizationID: 1, UserID: 1, Role: "owner"})
	repo.Add(models.OrganizationMember{OrganizationID: 2, UserID: 1, Role: "operator"})

	w := performAuthorizationRequest(setupAuthorizationRouter(repo), "2")

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	if w.Body.String() != `{"org_id":2,"role":"operator"}` {
		t.Fatalf("expected org 2 to be selected, got %s", w.Body.String())
	}
}

func TestAuthorizationMiddlewareRejectsDifferentOrgHeaderWithSingleMembership(t *testing.T) {
	repo := memory.NewOrganizationMemberRepository()
	repo.Add(models.OrganizationMember{OrganizationID: 1, UserID: 1, Role: "owner"})

	w := performAuthorizationRequest(setupAuthorizationRouter(repo), "2")

	if w.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", w.Code)
	}
}

func TestAuthorizationMiddlewareRequiresHeaderWithMultipleMemberships(t *testing.T) {
	repo := memory.NewOrganizationMemberRepository()
	repo.Add(models.OrganizationMember{OrganizationID: 1, UserID: 1, Role: "owner"})
	repo.Add(models.OrganizationMember{OrganizationID: 2, UserID: 1, Role: "operator"})

	w := performAuthorizationRequest(setupAuthorizationRouter(repo), "")

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestAuthorizationMiddlewareRequiresHeaderWithoutMemberships(t *testing.T) {
	repo := memory.NewOrganizationMemberRepository()

	w := performAuthorizationRequest(setupAuthorizationRouter(repo), "")

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}

func TestAuthorizationMiddlewareRejectsOrgHeaderWhenNotMember(t *testing.T) {
	repo := memory.NewOrganizationMemberRepository()
	repo.Add(models.OrganizationMember{OrganizationID: 1, UserID: 1, Role: "owner"})

	w := performAuthorizationRequest(setupAuthorizationRouter(repo), "99")

	if w.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", w.Code)
	}
}

func TestAuthorizationMiddlewareRejectsInvalidOrgHeader(t *testing.T) {
	repo := memory.NewOrganizationMemberRepository()
	repo.Add(models.OrganizationMember{OrganizationID: 1, UserID: 1, Role: "owner"})

	w := performAuthorizationRequest(setupAuthorizationRouter(repo), "abc")

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", w.Code)
	}
}
