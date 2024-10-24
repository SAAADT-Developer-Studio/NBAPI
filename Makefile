# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

build-windows:
	@echo "Building..."
	@go build -o main.exe cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go


# Create DB container
docker-run:
	@if docker compose up 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi


# Test the application
test:
	@echo "Testing..."
	@go test ./... -v


# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v


# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload

watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

watch-windows:
	@air -c ".air.windows.toml"

sqlc:
	@sqlc generate

seed-fast:
	@PGPASSWORD=postgres psql -h localhost -p 5432 -d postgres -U postgres -c "DROP SCHEMA postgres CASCADE;"
	@PGPASSWORD=postgres psql -h localhost -p 5432 -d postgres -U postgres -f data/backup.sql

seed-slow:
	@PGPASSWORD=postgres psql -h localhost -p 5432 -d postgres -U postgres -f data/schema.sql
	@go run scripts/parser.go

seed-prod:
	@echo "TODO"

pg-dump:
	@PGPASSWORD=postgres pg_dump -U postgres -h localhost -d postgres -f data/backup.sql

.PHONY: all build run test clean watch
