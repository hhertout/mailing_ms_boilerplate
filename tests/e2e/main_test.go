package e2e_test

import (
	"mailer_ms/internal/application/api/router"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestPingRoute(t *testing.T) {
	err := os.Setenv("GO_ENV", "development")
	if err != nil {
		t.Errorf("Failed to set env")
	}
	r := router.Serve()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)

	r.ServeHTTP(w, req)

	t.Run("Expect status code is ok", func(t *testing.T) {
		if w.Code != 200 {
			t.Errorf("Invalid request")
		}
	})
}
