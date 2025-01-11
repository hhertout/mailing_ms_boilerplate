package integration_test

import (
	"mailer_ms/internal/application/controllers"
	"os"
	"testing"
)

func TestNewController(t *testing.T) {
	err := os.Setenv("DB_URL", "../../data/mailerTest.db")
	if err != nil {
		t.Error("Failed to set db env")
	}

	a := controllers.NewApiController()
	if a == nil {
		t.Error("Repository is null")
	}
}
