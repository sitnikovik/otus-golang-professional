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


###> Build and run ###
DOCKER_IMG="calendar:develop"
GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

.PHONY: build
build:
	go build -v -o $(APP_BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

.PHONY: run
run: build
	$(BIN) -config ./.env

.PHONY: build-img
build-img:
	docker build \
		--build-arg=LDFLAGS="$(LDFLAGS)" \
		-t $(DOCKER_IMG) \
		-f build/Dockerfile .

.PHONY: run-img
run-img: build-img
	docker-compose build --no-cache
	docker-compose down
	docker-compose up --build
	docker run $(DOCKER_IMG)
###< Build and run ###

###> Test with Docker ###
.PHONY: test
test: start-test-db wait-for-test-db run-tests stop-test-db

.PHONY: start-test-db
start-test-db:
	docker-compose -f docker-compose-test.yaml up -d

.PHONY: wait-for-test-db
wait-for-test-db:
	@echo "Waiting for test database to be ready..."
	@until docker exec calendar_pg_test pg_isready -U $(PG_USER) -d $(PG_DATABASE); do \
		sleep 1; \
	done

.PHONY: run-tests
run-tests:
	@echo "Running tests..."
	PG_HOST=$(PG_HOST_TEST) \
	PG_PORT=$(PG_PORT_TEST) \
	PG_USER=$(PG_USER_TEST) \
	PG_PASSWORD=$(PG_PASSWORD_TEST) \
	PG_DATABASE=$(PG_DATABASE_TEST) \
	go test -race ./internal/...

.PHONY: stop-test-db
stop-test-db:
	docker-compose -f docker-compose-test.yaml down
###< Test with Docker ###

###> Local db migrations ###
LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=$(PG_HOST) port=$(PG_PORT) dbname=$(PG_DATABASE) user=$(PG_USER) password=$(PG_PASSWORD)"

.PHONY: local-migration-create
local-migration-status:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

.PHONY: local-migration-up
local-migration-up:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

.PHONY: local-migration-down
local-migration-down:
	goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v
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