package integration_test

import (
	"errors"
	"mailer_ms/src/repository"
	"os"
	"testing"
)

func TestWithoutError(t *testing.T) {
	// Setup
	err := os.Setenv("DB_URL", "./data/mailer_test.db")

	if err != nil {
		t.Error("Failed to set db env")
	}

	dbPool, err := connect()
	if err != nil {
		t.Error("Failed to connect to db")
	}

	r, err := repository.NewRepository(dbPool)
	if err != nil {
		t.Error("repository instantiation failed")
	}
	if r == nil {
		t.Error("Repository is null")
	}

	// testing method
	to := []string{"test@example.com", "another@example.com"}
	subject := "Test subject"

	err = r.SaveWithoutError(to, subject)
	if err != nil {
		t.Errorf("SaveWithoutError failed: %s", err)
	}
}

func TestSaveWithError(t *testing.T) {
	// Setup
	err := os.Setenv("DB_URL", "./data/mailer_test.db")

	if err != nil {
		t.Error("Failed to set db env")
	}

	dbPool, err := connect()
	if err != nil {
		t.Error("Failed to connect to db")
	}

	r, err := repository.NewRepository(dbPool)
	if err != nil {
		t.Error("repository instantiation failed")
	}
	if r == nil {
		t.Error("Repository is null")
	}

	// testing method
	to := []string{"test@example.com", "another@example.com"}
	subject := "Test subject"

	err = r.SaveWithError(to, subject, errors.New("failed"))
	if err != nil {
		t.Errorf("SaveWithoutError failed: %s", err)
	}
}

func TestListMail(t *testing.T) {
	// Setup
	if err := os.Setenv("DB_URL", "./data/mailer_test.db"); err != nil {
		t.Error("Failed to set db env")
	}

	dbPool, err := connect()
	if err != nil {
		t.Error("Failed to connect to db")
	}

	if _, err = dbPool.Exec(`DELETE FROM "mail"`); err != nil {
		t.Error("Failed to reset db")
	}

	r, err := repository.NewRepository(dbPool)
	if err != nil {
		t.Error("repository instantiation failed")
	}
	if r == nil {
		t.Error("Repository is null")
	}

	// >Insert some data with the 2 methods
	toWithoutError := []string{"test@example.com", "another@example.com"}
	subjectWithoutError := "Test subject without error"
	err = r.SaveWithoutError(toWithoutError, subjectWithoutError)
	if err != nil {
		t.Errorf("SaveWithoutError failed: %s", err)
	}
	toWithError := []string{"error@example.com", "fail@example.com"}
	subjectWithError := "Test subject with error"

	err = r.SaveWithError(toWithError, subjectWithError, errors.New("error"))
	if err != nil {
		t.Errorf("SaveWithError failed: %s", err)
	}

	// Test ListMail > must retrieve data inserted before
	mails, err := r.ListMail()
	if err != nil {
		t.Errorf("ListMail failed: %s", err)
	}

	if len(mails) < 2 {
		t.Error("Expected 2 min mails, got", len(mails))
	}

	mailWithErrIsContained := false
	mailWithoutErrIsContained := false

	for _, mail := range mails {
		if mail.Subject == "Test subject with error" {
			mailWithErrIsContained = true

			if mail.To != "error@example.com;fail@example.com;" {
				t.Errorf(" Expect error@example.com;fail@example.com;, got %v", mail.To)
			}

			if mail.Sent != false {
				t.Errorf("false, got %v", mail.Sent)
			}

		} else if mail.Subject == "Test subject without error" {
			mailWithoutErrIsContained = true

			if mail.To != "test@example.com;another@example.com;" {
				t.Errorf(" Expect test@example.com;another@example.com;, got %v", mail.To)
			}

			if mail.Sent != true {
				t.Errorf("true, got %v", mail.Sent)
			}
		}
	}

	if !mailWithErrIsContained || !mailWithoutErrIsContained {
		t.Error("Expected 2 mail inserted")
	}
}
