FROM golang:1.20.1-alpine AS builder

RUN apk update && \
  apk add make && \
  go install goa.design/goa/v3/cmd/goa@v3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN chmod +x /app/wait-for.sh

EXPOSE 8000

CMD ["make", "run"]