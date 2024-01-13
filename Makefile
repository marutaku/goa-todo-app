.PHONY: gen setup build run gen-migration migrate test lint clean

IS_ENVFILE_EXISTS =  $(shell ls | grep .env)
ifeq ($(IS_ENVFILE_EXISTS), .env)
	include .env
endif
DIST_PATH = ./dist
PRODUCTION_BUILD_FLAGS = -ldflags='-s -w'
ENV ?= development
# hostがlocalhostなのはローカルからの実行のみを想定しているため
POSTGRES_DSN = "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable&search_path=public"
DEV_URL = "docker://postgres/15/dev?search_path=public"

ifeq ($(ENV), production)
	BUILD_COMMAND="go build ${PRODUCTION_BUILD_FLAGS} -o ${DIST_PATH}/task ./cmd/task"
else
	BUILD_COMMAND="go build -o ${DIST_PATH}/task ./cmd/task"
endif



setup:
	go mod download

gen: 
	goa gen backend/design

build:
	mkdir -p ${DIST_PATH} && \
	eval ${BUILD_COMMAND}

build-cli:
	mkdir -p ${DIST_PATH} && \
	go build -o ${DIST_PATH}/task-cli ./cmd/task-cli

run:
	air -c .air.toml

test:
	go test -v -cover ./...

lint:
	golangci-lint run

clean:
	rm -rf ${DIST_PATH}
	air -c .air.toml

gen-migration:
	atlas migrate diff --env gorm --var dsn=$(POSTGRES_DSN)

validate-migration:
	atlas migrate validate --env gorm

migrate:
	atlas schema apply --env gorm --var dsn=$(POSTGRES_DSN)
