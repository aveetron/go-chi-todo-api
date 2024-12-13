dependencies-download:
	go mod tidy

build:
	go build -o main main.go

dev:
	go run main.go

db-up:
	docker compose up -d

migrate-up:
	go run migrations/main.go up

migrate-down:
	go run migrations/main.go down