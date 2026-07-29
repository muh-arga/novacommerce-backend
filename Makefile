.DEFAULT_GOAL := help

COMPOSE := docker compose

.PHONY: help dev up down build rebuild logs ps fmt vet test check clean

help: ## Show this help message
	@echo "NovaCommerce Backend"
	@echo ""
	@echo "Available commands:"
	@echo "  make dev        Start development environment"
	@echo "  make up         Start container in background"
	@echo "  make down       Stop and remove container"
	@echo "  make build      Build Docker image"
	@echo "  make rebuild    Rebuild and start containers"
	@echo "  make logs       Follow API container logs"
	@echo "  make ps         Show container status"
	@echo "  make fmt        Format Go source code"
	@echo "  make vet        Run Go static analysis"
	@echo "  make test       Run tests"
	@echo "  make check      Run format checking, vet, and test"
	@echo "  make clean      Stop containers and remove temporary files"

dev:
	$(COMPOSE) up --build

up:
	$(COMPOSE) up -d

down:
	$(COMPOSE) down

build:
	$(COMPOSE) build

rebuild:
	$(COMPOSE) down
	$(COMPOSE) up --build -d

logs:
	$(COMPOSE) logs -f api

ps:
	$(COMPOSE) ps

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

check: fmt vet test

clean:
	$(COMPOSE) down
	rm -rf tmp
