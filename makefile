build:
	@echo "Building the application..."
	@go build -o swif cmd/main.go

run:
	@echo "Running the application..."
	@go run ./cmd/main.go

up:
	@echo  "Docker compose up..."
	@docker compose up

git-sync:
	@echo "Git add... "
	@git add .
	@echo "git commit...."
	@git commit -m "Sync and committing...."

down:
	@echo "Docker compose down..."
	@docker compose down

remove:
	@docker rmi -r $(docker images -q)
.PHONY: run build

