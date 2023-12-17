package main

import (
	"fmt"
	"log"
	"mailer_ms/cmd/database"
	"mailer_ms/src/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("DOCKER") != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	db, err := database.Connect()
	migration := database.NewMigration(db)
	if err != nil {
		log.Println("Failed to connect to db")
	}
	if err = migration.Migrate(); err != nil {
		fmt.Printf("Failed to migrate db: %s", err)
	}

	r := router.Serve()
	log.Printf("📡 Server start on port %s \n", os.Getenv("PORT"))
	if err := r.Run(); err != nil {
		fmt.Println("Error on running server")
		fmt.Printf("Error: %s", err)

		return
	}
}
