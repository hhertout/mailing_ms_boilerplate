package main

import (
	"context"
	"log"
	"mailer_ms/internal/application/kafka/handler"
	"mailer_ms/migrations"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const (
	ENV_DOCKER         = "DOCKER"
	ENV_GO_ENV         = "GO_ENV"
	ENV_KAFKA_ENDPOINT = "KAFKA_ENDPOINT"
	ENV_KAFKA_GROUP_ID = "KAFKA_GROUP_ID"
)

var (
	TOPICS         = []string{"test-topic"}
	KAFKA_ENDPOINT = os.Getenv(ENV_KAFKA_ENDPOINT)
	KAFKA_GROUP_ID = os.Getenv(ENV_KAFKA_GROUP_ID)
)

func main() {
	if os.Getenv(ENV_DOCKER) != "true" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	logger, _ := zap.NewProduction()
	if os.Getenv(ENV_GO_ENV) == "development" {
		logger, _ = zap.NewDevelopment()
	}
	defer logger.Sync()

	m := migrations.NewMigration("/", logger)
	if err := m.MigrateAll(); err != nil {
		logger.Error(err.Error())
		return
	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": KAFKA_ENDPOINT,
		"group.id":          KAFKA_GROUP_ID,
		"auto.offset.reset": "latest",
	})
	if err != nil {
		logger.Fatal("Error while creating the consumer: ", zap.Error(err))
	}
	defer c.Close()

	err = c.SubscribeTopics(TOPICS, nil)
	if err != nil {
		logger.Fatal("Error while connecting to topic: ", zap.Error(err))
	}

	logger.Debug("Successfully connected to the topic, waiting for message...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigchan
		logger.Info("Termination signal received, closing consumer...")
		cancel()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := c.ReadMessage(-1)
			if err == nil {
				handler.HandleMessage(logger, msg)
			} else {
				logger.Sugar().Errorf("Error while reading the message: %v\n", err)
			}
		}
	}
}
