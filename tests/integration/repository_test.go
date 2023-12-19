package integration_test

import (
	"database/sql"
	"mailer_ms/cmd/database"
	"mailer_ms/src/repository"
	"os"
	"testing"
)

func connect() (*sql.DB, error) {
	err := os.Setenv("DB_URL", "../../data/mailerTest.db")
	if err != nil {
		return nil, err
	}

	db, errConnect := database.Connect()
	if errConnect != nil {
		return nil, err
	}

	migration := database.NewMigration(db, "/../..")
	if err = migration.Migrate(); err != nil {
		return nil, err
	}

	return db, nil
}

func TestNewRepository(t *testing.T) {
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
}
