include .env-example

.PHONY: build migrator dev local-env local-env-down test

build:
	go build -o bin/mnotes cmd/mnotes/main.go

migrate:
	@goose -dir migration postgres "host=${DB_HOST} \
			  user=${DB_USER} \
			  dbname=${DB_NAME} sslmode=disable \
			  password=${DB_PASSWORD}" up

dev: build
	./bin/mnotes
	
local-env:
	docker-compose up -d

local-env-down:
	docker-compose down

test:
	go test ./...

gen-key:
	openssl genrsa -out rsa_key 2048
	openssl rsa -in rsa_key -pubout > rsa_key.pub

deps-install:
	go get github.com/matryer/moq
