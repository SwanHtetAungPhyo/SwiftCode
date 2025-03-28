BINARY = swif
PORT = 8080
DATA_PATH = ./data/intern.csv


GREEN = \033[0;32m
YELLOW = \033[0;33m
NC = \033[0m

.PHONY: build run up down test ci_cd git-sync remove curl-load help

build:
	@echo "$(GREEN)Building the application...$(NC)"
	@go build -o $(BINARY) cmd/main.go

run-local:
	@echo "$(GREEN)Running the application on port $(PORT)...$(NC)"
	@PORT=$(PORT) DATA_PATH=$(DATA_PATH) go run cmd/main.go

up:
	@echo "$(GREEN)Starting Docker Compose...$(NC)"
	@docker compose up -d

down:
	@echo "$(YELLOW)Stopping Docker Compose...$(NC)"
	@docker compose down

git-sync:
	@echo "$(GREEN)Staging all changes...$(NC)"
	@git add .
	@echo "$(GREEN)Committing changes...$(NC)"
	@git commit -m "Sync: $(shell date +'%Y-%m-%d %H:%M:%S')"
	@echo "$(GREEN)Pushing changes...$(NC)"
	@git push origin main

doc:
	@pkgsite -http=:6060
ci_cd:
	@echo "$(GREEN)Running GitHub CI/CD with act...$(NC)"
	@act --container-architecture linux/amd64

test:
	@echo "$(GREEN)Running test cases...$(NC)"
	@go test -v ./swiftcode_test

remove:
	@echo "$(YELLOW)Removing unused Docker images...$(NC)"
	@docker rmi -f $(shell docker images -q)

curl-load:
	@echo "$(GREEN)Running load test with curl...$(NC)"
	@./http/curl.sh

help:
	@echo "$(GREEN)Available Commands:$(NC)"
	@echo "  make up        - Start the application with Docker Compose"
	@echo "  make down      - Stop and remove Docker containers"
	@echo "  make build     - Compile the Go application"
	@echo "  make run       - Run the application locally"
	@echo "  make test      - Run unit tests"
	@echo "  make ci_cd     - Run GitHub Actions locally using act"
	@echo "  make git-sync  - Commit and push changes"
	@echo "  make remove    - Remove unused Docker images"
	@echo "  make curl-load - Run a load test with curl"
	@echo "  make help      - Show this help message"
	@echo "Prometheus: http://localhost:9090"
	@echo "Grafana Login: admin/admin"
	@echo "Swagger API Docs: http://localhost:8080/swagger/index.html"
