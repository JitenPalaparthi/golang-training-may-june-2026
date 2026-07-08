package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	_ "net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPing(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/v1/api/public/ping", Ping)

	req := httptest.NewRequest(http.MethodGet, "/v1/api/public/ping", nil)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status  %d, got %d", http.StatusOK, rec.Code)
	}

	var resp map[string]string

	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("faild to unmarshal the response: %v", err)
	}

	if resp["ping"] != "pong" {
		t.Errorf("Expected ping=pong,got %v", resp["ping"])
	}

}
func TestHealth(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.GET("/v1/api/public/health", Health)

	req := httptest.NewRequest(http.MethodGet, "/v1/api/public/health", nil)

	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status  %d, got %d", http.StatusOK, rec.Code)
	}

	var resp map[string]string

	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("faild to unmarshal the response: %v", err)
	}

	if resp["health"] != "ok" {
		t.Errorf("Expected ping=pong,got %v", resp["ping"])
	}

}
