.PHONY: gen build run

gen: 
	goa gen backend/design

build:
	mkdir -p ./dist && \
	go build -o ./dist/task ./cmd/task && \
	go build -o ./dist/task-cli ./cmd/task-cli

run:
	air -c .air.toml