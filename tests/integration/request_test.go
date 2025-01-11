package integration_test

import (
	GoMailer "mailer_ms/pkg/mailer"
	"os"
	"testing"
)

func TestNewMailer(t *testing.T) {
	err := os.Setenv("SMTP_FROM", "toto@test.com")
	if err != nil {
		panic("failed to set env")
	}
	subject := "testing"
	to := "tata@test.com"
	mailer := GoMailer.NewRequest(subject, []string{to})

	t.Run("🧪 Expect new mailer not fail", func(t *testing.T) {
		if mailer == nil {
			t.Error("Expected mailer to be not nil")
		}
	})
}
