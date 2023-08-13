package GoMailer_test

import (
	GoMailer "mailer_ms/services"
	"os"
	"testing"
)

func TestNewMailer(t *testing.T) {
	os.Setenv("SMTP_FROM", "toto@test.com")
	subject := "testing"
	to := "tata@test.com"
	mailer := GoMailer.NewRequest(subject, []string{to})

	t.Run("🧪 Expect new mailer not fail", func(t *testing.T) {
		if mailer == nil {
			t.Error("Expected mailer to be not nil")
		}
	})
	t.Run("🧪 Expect From is correct", func(t *testing.T) {
		if mailer.From != "toto@test.com" {
			t.Errorf("Error on 'from' param got %v", mailer.From)
		}
	})
	t.Run("🧪 Expect To is correct", func(t *testing.T) {
		if mailer.From != "toto@test.com" {
			t.Errorf("Error on 'subject'. Expect %v, got %v", to, mailer.To)
		}
	})
	t.Run("🧪 Expect Subject is correct", func(t *testing.T) {
		if mailer.Subject != "testing" {
			t.Errorf("Error on 'subject'. Expect %v, got %v", subject, mailer.Subject)
		}
	})
}

func TestSetPlainTextBody(t *testing.T) {
	os.Setenv("SMTP_FROM", "toto@test.com")
	subject := "testing"
	to := "tata@test.com"
	mailer := GoMailer.NewRequest(subject, []string{to})
	mailer.SetPlainTextBody("Hello World")

	t.Run("🧪 Expect new mailer not fail", func(t *testing.T) {
		if mailer == nil {
			t.Error("Expected mailer to be not nil")
		}
	})
	t.Run("🧪 Expect Body is correct", func(t *testing.T) {
		if mailer.Body != "Hello World" {
			t.Errorf("Error on 'Body'. Expect %v, got %v", "Hello World", mailer.Body)
		}
	})
}
