build:
	@echo "Building the application..."
	@go build -o swif cmd/main.go

run:
	@echo "Running the application..."
	@PORT=8080 DATA_PATH=./data/intern.csv go run main.go

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

