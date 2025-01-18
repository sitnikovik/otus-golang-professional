include .env

LOCAL_BIN:=$(CURDIR)/bin
APP_BIN := "$(LOCAL_BIN)/calendar"

.PHONY: version
version: build
	$(BIN) version

.PHONY: install-lint-deps
install-lint-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: lint
lint: install-lint-deps
	golangci-lint run --config=.golangci.yml

.PHONY: test
test: unit-test

.PHONY: unit-test
unit-test:
	@echo "Running unit tests..."
	go test -race -tags=unit ./...

.PHONY: mocks
mocks:
	@echo "Generating mocks..."
	go install github.com/vektra/mockery/v2/...@latest
	mockery

###> Build and run ###
DOCKER_IMG="calendar:develop"
DOCKER_SCHEDULER_IMG="scheduler:develop"
DOCKER_SENDER_IMG="sender:develop"
API_DOCKER_COMPOSE_FILE="build/calendar/docker-compose.yaml"
API_TEST_DOCKER_COMPOSE_FILE="docker-compose-test.yaml"
SCHEDULER_DOCKER_COMPOSE_FILE="build/scheduler/docker-compose.yaml"
GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

.PHONY: build
build: build-calendar build-scheduler

.PHONY: build-calendar
build-calendar:
	go build -v -o $(APP_BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

.PHONY: build-scheduler
build-scheduler:
	go build -v -o $(APP_BIN) -ldflags "$(LDFLAGS)" ./cmd/scheduler

.PHONY: build-sender
build-sender:
	go build -v -o $(APP_BIN) -ldflags "$(LDFLAGS)" ./cmd/sender

.PHONY: run
run: build
	$(BIN) -config ./.env

.PHONY: build-calendar-img
build-calendar-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/calendar/Dockerfile .

.PHONY: build-scheduler-img
build-scheduler-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_SCHEDULER_IMG) \
		-f build/scheduler/Dockerfile .

.PHONY: build-sender-img
build-sender-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_SENDER_IMG) \
		-f build/sender/Dockerfile .

.PHONY: run-calendar-img
run-calendar-img: build-calendar-img
	docker-compose build calendar_app --no-cache
	docker-compose down calendar_app
	docker-compose up calendar_app --build
	docker run $(DOCKER_IMG)
	
.PHONY: run-scheduler-img
run-scheduler-img: build-scheduler-img
	docker-compose build scheduler_app --no-cache
	docker-compose down scheduler_app
	docker-compose up scheduler_app --build
	docker run $(DOCKER_SCHEDULER_IMG)

.PHONY: run-sender-img
run-sender-img: build-sender-img
	docker-compose build sender_app --no-cache
	docker-compose down sender_app
	docker-compose up sender_app --build
	docker run $(DOCKER_SENDER_IMG)
###< Build and run ###

###> Test with Docker ###
.PHONY: integration-test
integration-test: start-test-db wait-for-test-db run-integration-tests stop-test-db

.PHONY: start-test-db
start-test-db:
	docker-compose -f $(API_TEST_DOCKER_COMPOSE_FILE) up -d

.PHONY: wait-for-test-db
wait-for-test-db: local-migration-up-test 
	@echo "Waiting for test database to be ready..."
	@until docker exec calendar_pg_test pg_isready -U $(PG_USER) -d $(PG_DATABASE); do \
		sleep 1; \
	done

.PHONY: run-integration-tests
run-integration-tests:
	@echo "Running integration tests..."
	PG_HOST=$(PG_HOST_TEST) \
	PG_PORT=$(PG_PORT_TEST) \
	PG_USER=$(PG_USER_TEST) \
	PG_PASSWORD=$(PG_PASSWORD_TEST) \
	PG_DATABASE=$(PG_DATABASE_TEST) \
	go test -race -tags=integration ./...

.PHONY: stop-test-db
stop-test-db: local-migration-down-test
	docker-compose -f $(API_TEST_DOCKER_COMPOSE_FILE) down
###< Test with Docker ###

###> Local db migrations ###
LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=$(PG_HOST_LOCAL) port=$(PG_PORT_LOCAL) dbname=$(PG_DATABASE) user=$(PG_USER) password=$(PG_PASSWORD)"
LOCAL_MIGRATION_DSN_TEST="host=$(PG_HOST_TEST) port=$(PG_PORT_TEST) dbname=$(PG_DATABASE_TEST) user=$(PG_USER_TEST) password=$(PG_PASSWORD_TEST)"

.PHONY: local-migration-create
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

.PHONY: local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONY: local-migration-up-test
local-migration-up-test:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN_TEST} up -v

.PHONY: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v
	.PHONY: local-migration-down

local-migration-down-test:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN_TEST} down -v
###< Local db migrations ###

###> grpc ###
install-go-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

generate: install-go-deps
	rm -r pkg/calendar/v1
	mkdir -p pkg/calendar/v1
	protoc --proto_path api/calendar/v1 \
	--go_out=pkg/calendar/v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/calendar/v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/calendar/v1/EventService.proto
###< grpc ###