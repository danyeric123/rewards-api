.PHONY: build run clean

help: ## Show this help message
	@echo "Usage: make [command]"
	@echo
	@echo "Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

check_env: ## Check if .env file exists
	@test -f .env || { \
			echo "\033[31mError: .env file not found.\033[0m"; \
			echo "\033[33mPlease create a .env file with the following variables:\033[0m"; \
			echo "\033[32m  POSTGRES_DB, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOST\033[0m"; \
			exit 1; \
	}

check_network: ## Check if docker network exists and if not, create it
	@docker network inspect rewards_network &>/dev/null || docker network create rewards_network

build: check_env check_network ## Build docker containers
	docker compose build

run: check_env check_network ## Run docker containers, rebuild if needed
	docker compose up --build

clean: ## Stop and remove docker containers
	docker compose down

fmt: ## Format code
	go fmt ./...

lint: ## Lint code
	golangci-lint run -v ./...

test: ## Run tests
	go test ./...

test_package: ## Run tests in package
	go test -v ./$(PACKAGE)
	