.PHONY: gen build

gen: 
	goa gen backend/design
	rm -rf ./cmd
	goa example backend/design

build:
	mkdir -p ./dist && \
	go build -o ./dist/todo ./cmd/todo && \
	go build -o ./dist/todo-cli ./cmd/todo-cli