.PHONY: gen build

gen: 
	goa gen backend/design
	rm -rf ./cmd
	goa example backend/design

build:
	mkdir -p ./dist && \
	go build -o ./dist/task ./cmd/task && \
	go build -o ./dist/task-cli ./cmd/task-cli