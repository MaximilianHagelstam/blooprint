# Build the application
build:	
	@go build -o bin/main cmd/main.go

# Run the application
run:
	@go run cmd/main.go

# Test the application
test:
	@go test ./... -v -cover

# Create DB container
db-up:
	@docker compose up

# Shutdown DB container
db-down:
	@docker compose down

# Clean the binary
clean:
	@rm -rf bin

# Live reload
watch:
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "installing air..."; \
		go install github.com/cosmtrek/air@latest; \
		air; \
	fi

# Lint
lint:
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "installing golangci-lint..."; \
		go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest; \
		golangci-lint run; \
	fi

.PHONY: build run test clean lint
