package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/atranna/atranna-api/src/internal/config"
	"github.com/atranna/atranna-api/src/internal/models"
	"github.com/atranna/atranna-api/src/internal/repository/memory"

	"github.com/gin-gonic/gin"
)

func TestGetUsersReturnsAllUsers(t *testing.T) {
	userRepo := memory.NewUserRepository()
	orgMemberRepo := memory.NewOrganizationMemberRepository()

	alice, _ := userRepo.Add(models.User{Username: "alice", Email: "alice@example.com", DisplayName: "Alice"})
	userRepo.Add(models.User{Username: "bob", Email: "bob@example.com", DisplayName: "Bob"})
	userRepo.Add(models.User{Username: "carol"})

	handler := NewHandler(userRepo, orgMemberRepo)
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("user_id", alice.ID)
		c.Next()
	})
	router.GET("/users", handler.GetUsers)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var response []struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(response) != 3 {
		t.Fatalf("expected 3 users, got %d", len(response))
	}

	seen := map[int]bool{}
	for _, u := range response {
		seen[u.ID] = true
	}
	if !seen[alice.ID] {
		t.Fatalf("expected alice in response, got %+v", response)
	}

	if strings.Contains(w.Body.String(), "email") || strings.Contains(w.Body.String(), "display_name") {
		t.Fatalf("response must not contain email or display_name: %s", w.Body.String())
	}
}

func TestGetUserOnlyExposesSharedOrganizations(t *testing.T) {
	userRepo := memory.NewUserRepository()
	orgMemberRepo := memory.NewOrganizationMemberRepository()

	alice, _ := userRepo.Add(models.User{Username: "alice", Email: "alice@example.com"})
	bob, _ := userRepo.Add(models.User{Username: "bob"})

	orgMemberRepo.Add(models.OrganizationMember{OrganizationID: 1, UserID: alice.ID, Role: "owner"})
	orgMemberRepo.Add(models.OrganizationMember{OrganizationID: 1, UserID: bob.ID, Role: "viewer"})
	orgMemberRepo.Add(models.OrganizationMember{OrganizationID: 2, UserID: bob.ID, Role: "viewer"})

	handler := NewHandler(userRepo, orgMemberRepo)
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("user_id", alice.ID)
		c.Next()
	})
	router.GET("/users/:id", handler.GetUser)

	req := httptest.NewRequest(http.MethodGet, "/users/2", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var response struct {
		ID   int   `json:"id"`
		Orgs []int `json:"orgs"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.ID != bob.ID {
		t.Fatalf("expected user %d, got %d", bob.ID, response.ID)
	}
	if len(response.Orgs) != 1 || response.Orgs[0] != 1 {
		t.Fatalf("expected only the shared organization [1], got %+v", response.Orgs)
	}
}

func TestDeleteUserOnlyAllowsSelf(t *testing.T) {
	userRepo := memory.NewUserRepository()
	orgMemberRepo := memory.NewOrganizationMemberRepository()

	alice, _ := userRepo.Add(models.User{Username: "alice"})
	bob, _ := userRepo.Add(models.User{Username: "bob"})

	handler := NewHandler(userRepo, orgMemberRepo)
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("user_id", alice.ID)
		c.Next()
	})
	router.DELETE("/users/:id", handler.Delete)

	req := httptest.NewRequest(http.MethodDelete, "/users/2", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Fatalf("expected 403 when deleting another user, got %d", w.Code)
	}

	if _, found := userRepo.GetByID(bob.ID); !found {
		t.Fatalf("expected bob to still exist after forbidden delete")
	}

	req = httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204 when deleting self, got %d", w.Code)
	}

	if _, found := userRepo.GetByID(alice.ID); found {
		t.Fatalf("expected alice to be deleted")
	}
}

func TestAddUserCanBeDisabledByConfig(t *testing.T) {
	original := config.Current.Auth.UserCreationEnabled
	config.Current.Auth.UserCreationEnabled = false
	defer func() { config.Current.Auth.UserCreationEnabled = original }()

	handler := NewHandler(memory.NewUserRepository(), memory.NewOrganizationMemberRepository())
	router := gin.Default()
	router.POST("/users", handler.Add)

	body, _ := json.Marshal(map[string]any{
		"username": "alice",
		"password": "secret",
	})
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusForbidden {
		t.Fatalf("expected 403 when user creation is disabled, got %d", w.Code)
	}
}

func TestAddUserWhenCreationEnabled(t *testing.T) {
	original := config.Current.Auth.UserCreationEnabled
	config.Current.Auth.UserCreationEnabled = true
	defer func() { config.Current.Auth.UserCreationEnabled = original }()

	handler := NewHandler(memory.NewUserRepository(), memory.NewOrganizationMemberRepository())
	router := gin.Default()
	router.POST("/users", handler.Add)

	body, _ := json.Marshal(map[string]any{
		"username": "alice",
		"password": "secret",
	})
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 when user creation is enabled, got %d", w.Code)
	}
}
