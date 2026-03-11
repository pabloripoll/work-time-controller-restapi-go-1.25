# -------------------------------------------------------------------------------------------------
#  Repository Helper to catch parent Makefile recipes
# -------------------------------------------------------------------------------------------------
%:
	@$(MAKE) -C .. $@

# -------------------------------------------------------------------------------------------------
#  Repository Helper for application recipes
# -------------------------------------------------------------------------------------------------
.PHONY: go-help go-run go-build go-test go-migrate-up go-migrate-down

APP="cmd/api/main.go"
APP_BUILD="bin/api $(APP)"
go-help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

go-run: ## Run the application
	go run $(APP)

go-build: ## Build the application
	go build -o $(APP_BUILD)

go-test: ## Run tests
	go test -v ./...

go-clean: ## Clean build artifacts
	rm -rf bin/
	rm -rf tmp/

go-install: ## Install dependencies
	go mod download
	go mod tidy

go-migrate-up: ## Run database migrations
	psql -h localhost -U postgres -d worktc -f migrations/001_initial_schema.sql

go-migrate-down: ## Rollback database migrations
	@echo "Manual rollback required - drop tables in reverse order"

go-dev: ## Run with hot reload (requires air)
	air

go-lint: ## Run linter
	golangci-lint run

go-format: ## Format code
	go fmt ./...

go-deps: ## Show dependencies
	go list -m all

.DEFAULT_GOAL := help
