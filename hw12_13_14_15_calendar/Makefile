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
API_TEST_DOCKER_COMPOSE_FILE="docker-compose-test.yaml"
GIT_HASH := $(shell git log --format="%h" -n 1)
LDFLAGS := -X main.release="develop" -X main.buildDate=$(shell date -u +%Y-%m-%dT%H:%M:%S) -X main.gitHash=$(GIT_HASH)

.PHONY: build
build: build-calendar-img build-scheduler-img build-sender-img

.PHONY: build-calendar
build-calendar:
	eval $(minikube docker-env)
	go build -v -o $(APP_BIN) -ldflags "$(LDFLAGS)" ./cmd/calendar

.PHONY: build-scheduler
build-scheduler:
	eval $(minikube docker-env)
	go build -v -o $(APP_BIN) -ldflags "$(LDFLAGS)" ./cmd/scheduler

.PHONY: build-sender
build-sender:
	eval $(minikube docker-env)
	go build -v -o $(APP_BIN) -ldflags "$(LDFLAGS)" ./cmd/sender

.PHONY: run
run: build
	$(BIN) -config ./.env

.PHONY: up
up: build
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

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
	docker-compose up calendar_app -d --build
	
.PHONY: down-calendar-img
down-calendar-img: 
	docker-compose down calendar_app

.PHONY: down-calendar-db
down-calendar-db: 
	docker-compose down calendar_pg
	
.PHONY: run-scheduler-img
run-scheduler-img: build-scheduler-img
	docker-compose build scheduler_app --no-cache
	docker-compose down scheduler_app
	docker-compose up scheduler_app -d --build
	docker run $(DOCKER_SCHEDULER_IMG)

.PHONY: down-scheduler-img
down-scheduler-img: 
	docker-compose down scheduler_app

.PHONY: run-sender-img
run-sender-img: build-sender-img
	docker-compose build sender_app --no-cache
	docker-compose down sender_app
	docker-compose up sender_app -d --build
	docker run $(DOCKER_SENDER_IMG)

.PHONY: down-sender-img
down-sender-img: 
	docker-compose down sender_app

.PHONY: down-rabbitmq
down-rabbitmq: 
	docker-compose down rabbitmq

###< Build and run ###

###> Test with Docker ###
.PHONY: integration-tests
integration-tests: start-test-db wait-for-test-db
	@$(MAKE) run-integration-tests && \
	$(MAKE) stop-test-db && \
	echo "✅ Integration tests passed" && exit 0 || \
	($(MAKE) stop-test-db && echo "❌ Integration tests failed" && exit 1)

.PHONY: start-test-db
start-test-db:
	. ./configs/setenv.sh .env.test && \
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
	go test -race -tags=integration ./...

.PHONY: stop-test-db
stop-test-db: local-migration-down-test
	docker-compose -f $(API_TEST_DOCKER_COMPOSE_FILE) down
###< Test with Docker ###

###> Local db migrations ###
.PHONY: local-migration-status
local-migration-status:
	. ./configs/setenv.sh .env && \
	./migrations/postgres/status.sh

.PHONY: local-migration-up
local-migration-up:
	. ./configs/setenv.sh .env && \
	./migrations/postgres/up.sh -c

.PHONY: local-migration-up-test
local-migration-up-test:
	. ./configs/setenv.sh .env.test && \
	./migrations/postgres/up.sh

.PHONY: local-migration-down
local-migration-down:
	. ./configs/setenv.sh .env && \
	./migrations/postgres/down.sh

.PHONY: local-migration-down-test
local-migration-down-test:
	. ./configs/setenv.sh .env.test && \
	./migrations/postgres/down.sh
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

###> k8s ###
.PHONY: k8s
k8s:
	minikube start
	minikube addons enable ingress
	eval $$(minikube docker-env)
	docker build -t calendar:develop -f build/calendar/Dockerfile .
	docker build -t scheduler:develop -f build/scheduler/Dockerfile .
	docker build -t sender:develop -f build/sender/Dockerfile .
	helm uninstall calendar --ignore-not-found
	helm install calendar ./helm/calendar-chart
	@echo "Add the following to /etc/hosts:"
	@echo "$$(minikube ip)  calendar.local"
	@echo "Then open http://calendar.local in your browser."

k8s-purge:
	kubectl delete svc calendar-pg --ignore-not-found
	kubectl delete svc calendar-service --ignore-not-found
	kubectl delete svc rabbitmq --ignore-not-found
	kubectl delete svc scheduler-service --ignore-not-found
	kubectl delete svc sender-service --ignore-not-found
	kubectl delete deployment calendar-app --ignore-not-found
	kubectl delete deployment postgres --ignore-not-found
	kubectl delete deployment rabbitmq --ignore-not-found
	kubectl delete deployment scheduler-app --ignore-not-found
	kubectl delete deployment sender-app --ignore-not-found
	kubectl delete ingress calendar-ingress --ignore-not-found