.PHONY: build run test clean docker-build docker-run migrate-up migrate-down

# Variáveis
APP_NAME=agencia-viagens
DOCKER_COMPOSE=docker-compose

# Comandos Go
build:
	go build -o bin/$(APP_NAME) ./cmd/api

run:
	go run ./cmd/api

test:
	go test -v -cover ./...

clean:
	rm -rf bin/
	rm -rf tmp/

# Comandos Docker
docker-build:
	$(DOCKER_COMPOSE) build

docker-run:
	$(DOCKER_COMPOSE) up -d

docker-stop:
	$(DOCKER_COMPOSE) down

# Comandos de Migração
migrate-up:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/agencia_viagens?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgresql://postgres:postgres@localhost:5432/agencia_viagens?sslmode=disable" down

# Comandos de Desenvolvimento
dev:
	air -c .air.toml

# Comandos de Dependências
deps:
	go mod download
	go mod tidy

# Comandos de Lint
lint:
	golangci-lint run

# Comandos de Railway
railway-up:
	railway up

railway-logs:
	railway logs 