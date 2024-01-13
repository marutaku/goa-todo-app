.PHONY: gen setup build run gen-migration migrate

include .env

# hostがlocalhostなのはローカルからの実行のみを想定しているため
POSTGRES_DSN = "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable&search_path=public"
DEV_URL = "docker://postgres/15/dev?search_path=public"

setup:
	go mod download

gen: 
	goa gen backend/design

build:
	mkdir -p ./dist && \
	go build -o ./dist/task ./cmd/task && \
	go build -o ./dist/task-cli ./cmd/task-cli

run:
	air -c .air.toml

gen-migration:
	atlas migrate diff --env gorm --var dsn=$(POSTGRES_DSN)

validate-migration:
	atlas migrate validate --env gorm

migrate:
	atlas schema apply --env gorm --var dsn=$(POSTGRES_DSN)