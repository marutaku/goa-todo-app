.PHONY: gen build run gen-migration migrate

include .env

POSTGRES_DSN = "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable"
DEV_URL = "docker://postgres/15/dev?search_path=public"

gen: 
	goa gen backend/design

build:
	mkdir -p ./dist && \
	go build -o ./dist/task ./cmd/task && \
	go build -o ./dist/task-cli ./cmd/task-cli

run:
	air -c .air.toml

gen-migration:
	atlas migrate diff --env gorm  

validate-migration:
	atlas migrate validate --dir "./migrations" --dev-url $(POSTGRES_DSN)

migrate:
	atlas migrate apply --dir "./migrations" --url $(DEV_URL) --env gorm