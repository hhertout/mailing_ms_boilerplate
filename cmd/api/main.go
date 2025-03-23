package main

import (
	"fmt"
	"log"
	"mailer_ms/internal/application/router"
	"mailer_ms/migrations"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	if os.Getenv("DOCKER") != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	logger, _ := zap.NewProduction()
	if os.Getenv("GO_ENV") == "development" {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()

	m := migrations.NewMigration("/", logger)
	if err := m.MigrateAll(); err != nil {
		log.Println(err)
		return
	}

	r := router.Serve()
	log.Printf("📡 Server start on port %s \n", os.Getenv("PORT"))
	if err := r.Run(); err != nil {
		fmt.Println("Error on running server")
		fmt.Printf("Error: %s", err)

		return
	}
}
