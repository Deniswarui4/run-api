.PHONY: help build run test clean docker-build docker-up docker-down migrate dev

# Variables
APP_NAME=event-ticketing-api
DOCKER_IMAGE=event-ticketing-api:latest
GO_FILES=$(shell find . -name '*.go' -type f)

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

build: ## Build the application
	@echo "Building $(APP_NAME)..."
	@go build -o bin/$(APP_NAME) cmd/api/main.go
	@echo "Build complete: bin/$(APP_NAME)"

run: ## Run the application
	@echo "Running $(APP_NAME)..."
	@go run cmd/api/main.go

dev: ## Run the application in development mode with hot reload (requires air)
	@echo "Starting development server..."
	@air

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE) .
	@echo "Docker image built: $(DOCKER_IMAGE)"

docker-up: ## Start Docker containers
	@echo "Starting Docker containers..."
	@docker-compose up -d
	@echo "Containers started. API available at http://localhost:8080"

docker-down: ## Stop Docker containers
	@echo "Stopping Docker containers..."
	@docker-compose down
	@echo "Containers stopped"

docker-logs: ## View Docker logs
	@docker-compose logs -f api

migrate: ## Run database migrations (requires running DB)
	@echo "Running migrations..."
	@go run cmd/api/main.go migrate
	@echo "Migrations complete"

deps: ## Download dependencies
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies downloaded"

lint: ## Run linter (requires golangci-lint)
	@echo "Running linter..."
	@golangci-lint run
	@echo "Linting complete"

format: ## Format code
	@echo "Formatting code..."
	@go fmt ./...
	@echo "Formatting complete"

install-tools: ## Install development tools
	@echo "Installing development tools..."
	@go install github.com/cosmtrek/air@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@echo "Tools installed"

.DEFAULT_GOAL := help
