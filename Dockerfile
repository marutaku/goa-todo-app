FROM golang:1.20.1-alpine AS builder

RUN apk update && \
  apk add make && \
  go install goa.design/goa/v3/cmd/goa@v3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/dist/todo .
COPY --from=builder /app/dist/todo-cli .

COPY .env .env

EXPOSE 8000

CMD [ "/app/todo" ]