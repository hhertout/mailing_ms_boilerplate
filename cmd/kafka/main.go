package main

import (
	"log"
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
}
