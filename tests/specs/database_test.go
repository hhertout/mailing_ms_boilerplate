package specs

import (
	"fmt"
	"mailer_ms/cmd/database"
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
		fmt.Println("Failed to connect to db")
		fmt.Println(err)
		t.Error("Failed to connect to db")
	}

	migration := database.NewMigration(db, "/../..")
	if errMigration := migration.Migrate(); errMigration != nil {
		fmt.Println(errMigration)
		t.Errorf("Failed to migrate db: %s", errMigration)
	}
}
