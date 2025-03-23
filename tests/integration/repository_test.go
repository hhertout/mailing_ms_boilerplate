package integration_test

import (
	"database/sql"
	"mailer_ms/internal/infra/database"
	"mailer_ms/internal/infra/repository"
	"mailer_ms/migrations"
	"os"
	"testing"

	"go.uber.org/zap"
)

func connect() (*sql.DB, error) {
	logger, _ := zap.NewProduction()
	if os.Getenv("GO_ENV") == "development" {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()

	err := os.Setenv("DB_URL", "../../data/mailerTest.db")
	if err != nil {
		return nil, err
	}

	db, errConnect := database.Connect()
	if errConnect != nil {
		return nil, err
	}

	migration := migrations.NewMigration("/../../", logger)
	if err = migration.MigrateAll(); err != nil {
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
