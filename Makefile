#!make
-include .env

build: # Build the application	
	@go build -o bin/main cmd/main.go

run: # Run the application
	@go run cmd/main.go

test: # Run tests
	@go test ./... -v

test-coverage: # Run tests with coverage
	@go test ./... -v -cover

db-start: # Start DB container
	@docker compose up

db-stop: # Stop DB container
	@docker compose down

db-status: # Show DB migration status
	@cd migrations; goose postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_DATABASE} sslmode=disable" status

clean: # Clean the binary
	@go clean
	@rm -rf bin

watch: # Live reload
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "installing air..."; \
		go install github.com/cosmtrek/air@latest; \
		air; \
	fi

lint: # Lint
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "installing golangci-lint..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run ./...; \
	fi

docker-build: # Build docker image
	@docker build . -t gostarter

docker-run: # Run docker image
	@docker run gostarter:latest

migrations-up: # Run all up migrations
	@cd migrations; goose postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_DATABASE} sslmode=disable" up

migrations-down: # Run all down migrations
	@cd migrations; goose postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_DATABASE} sslmode=disable" down

help: # Print help
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: build run test test-coverage db-start db-stop db-status clean lint docker-build docker-run help
