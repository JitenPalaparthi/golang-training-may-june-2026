package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	calledNext = false
)

func TestAuth(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name       string
		username   string
		password   string
		wantStatus int
		wantNext   bool
	}{
		{
			name:       "valid credentials",
			username:   "admin",
			password:   "admin123",
			wantStatus: http.StatusOK,
			wantNext:   true,
		},

		{
			name:       "invalid username",
			username:   "wrong",
			password:   "admin123",
			wantStatus: http.StatusUnauthorized,
			wantNext:   false,
		},
		{
			name:       "invalid password ",
			username:   "wrong",
			password:   "wrongpassword",
			wantStatus: http.StatusUnauthorized,
			wantNext:   false,
		},
		{
			name:       "missing headers",
			username:   "",
			password:   "",
			wantStatus: http.StatusUnauthorized,
			wantNext:   false,
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			rounter := gin.New()
			rounter.Use(Auth)

			rounter.GET("/test", test)

			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			if tt.username != "" {
				req.Header.Set("username", tt.username)
			}
			if tt.password != "" {
				req.Header.Set("password", tt.password)
			}

			rec := httptest.NewRecorder()

			rounter.ServeHTTP(rec, req)
			t.Logf("expected status :%d , got %d\n", tt.wantStatus, rec.Code)
			if rec.Code != tt.wantStatus {
				t.Logf("expected status :%d , got %d\n", tt.wantStatus, rec.Code)
				t.Errorf("expected status :%d , got %d", tt.wantStatus, rec.Code)
			}
			t.Logf("expected next called  :%v , got %v\n", tt.wantNext, calledNext)
			if calledNext != tt.wantNext {
				t.Logf("expected next called  :%v , got %v\n", tt.wantNext, calledNext)
				t.Errorf("expected next called  :%v , got %v", tt.wantNext, calledNext)
			}

		})

	}

}

func test(ctx *gin.Context) {
	calledNext = true
	ctx.String(http.StatusOK, "success")
}
