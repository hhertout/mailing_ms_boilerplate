package main

import (
	"log"
	"mailer_ms/internal/application/api/router"
	"mailer_ms/migrations"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const (
	ENV_DOCKER = "DOCKER"
	ENV_GO_ENV = "GO_ENV"
	ENV_PORT   = "PORT"
)

func main() {
	// Load environment variables from .env file if not running in Docker
	if os.Getenv(ENV_DOCKER) != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Initialize the logger
	logger, _ := zap.NewProduction()
	if os.Getenv(ENV_GO_ENV) == "development" {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()

	// Run migrations
	m := migrations.NewMigration("/", logger)
	if err := m.MigrateAll(); err != nil {
		logger.Sugar().Errorf("Migration error: %v", err)
		return
	}

	// Start the server
	r := router.Serve()
	port := os.Getenv(ENV_PORT)
	logger.Sugar().Infof("📡 Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		logger.Sugar().Errorf("Error running server: %v", err)
		return
	}
}
