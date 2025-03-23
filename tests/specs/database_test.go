package specs

import (
	"fmt"
	"mailer_ms/migrations"
	"os"
	"testing"

	"go.uber.org/zap"
)

func TestDatabase(t *testing.T) {
	logger, _ := zap.NewProduction()
	if os.Getenv("GO_ENV") == "development" {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()

	err := os.Setenv("DB_URL", "../../data/mailerTest.db")
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to set DB_URL")
	}

	migration := migrations.NewMigration("/../../", logger)
	if errMigration := migration.MigrateAll(); errMigration != nil {
		fmt.Println(errMigration)
		t.Errorf("Failed to migrate db: %s", errMigration)
	}
}
