package tests_test

import (
	"mailer_ms/src/router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
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
