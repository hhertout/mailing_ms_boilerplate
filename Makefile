# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

build-cli:
	@go build -o tmp/cli cmd/cli/main.go

build-kafka:
	@go build -o tmp/kafka cmd/kafka/main.go

# reset sqlite db
db-reset:
	@docker compose down && rm ./data/mailer.db && docker compose up -d

# Create DB container
dc-up:
	@echo "Starting your application..."
	@docker compose up -d

# Shutdown DB container
dc-down:
	@echo "Closing..."
	@docker compose down

# Test the application
test:
	@echo "Testing..."
	@go test ./tests/... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@go mod tidy

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
