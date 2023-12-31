package specs

import (
	"fmt"
	"mailer_ms/migrations"
	"mailer_ms/src/database"
	"os"
	"testing"
)

func TestDatabase(t *testing.T) {
	err := os.Setenv("DB_URL", "../../data/mailerTest.db")
	if err != nil {
		fmt.Println(err)
		t.Error("Failed to set DB_URL")
	}

	db, errConnect := database.Connect()
	if errConnect != nil {
		t.Error("Failed to connect to db")
	}

	migration := migrations.NewMigration(db, "/../../")
	if errMigration := migration.MigrateAll(); errMigration != nil {
		fmt.Println(errMigration)
		t.Errorf("Failed to migrate db: %s", errMigration)
	}
}
