include .env-example

.PHONY: build migrator

build:
	go build -o bin/mnotes cmd/mnotes/main.go

migrate:
	goose postgres "host=${DB_HOST} \
			  user=${DB_USER} \
			  dbname=${DB_NAME} sslmode=disable \
			  password=${DB_PASSWORD}" \
			  -dir migration

dev: build
	./bin/mnotes
	
local-env:
	docker-compose up -d

local-env-down:
	docker-compose down

test:
	go test ./...
