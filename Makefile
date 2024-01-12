.PHONY: gen build run test lint clean

DIST_PATH = ./dist
PRODUCTION_BUILD_FLAGS = -ldflags='-s -w'
ENV ?= development

ifeq ($(ENV), production)
	BUILD_COMMAND="go build ${PRODUCTION_BUILD_FLAGS} -o ${DIST_PATH}/task ./cmd/task"
else
	BUILD_COMMAND="go build -o ${DIST_PATH}/task ./cmd/task"
endif

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
