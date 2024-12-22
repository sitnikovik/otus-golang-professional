include .env

BIN := "./bin/calendar"
DOCKER_IMG="calendar:develop"

GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

build:
	go build -v -o $(BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

run: build
	$(BIN) -config ./.env

build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .
	
	

run-img: build-img
	docker-compose build --no-cache
	docker-compose down
	docker-compose up --build
	docker run $(DOCKER_IMG)

version: build
	$(BIN) version

test:
	go test -race ./internal/...

install-lint-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

lint: install-lint-deps
	golangci-lint run --config=.golangci.yml

.PHONY: build run build-img run-img version test lint

LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=$(PG_HOST) port=$(PG_PORT) dbname=$(PG_DATABASE) user=$(PG_USER) password=$(PG_PASSWORD)"

## Local db migrations
.PHONY: local-migration-create
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

.PHONY: local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONY: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

