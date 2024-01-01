FROM golang:1.20.1-alpine AS builder

RUN apk update && \
  apk add make && \
  go install goa.design/goa/v3/cmd/goa@v3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./dist/task ./cmd/task

FROM alpine:3.13

WORKDIR /app

COPY --from=builder /app/dist/task .
COPY --from=builder /app/dist/task-cli .
COPY --from=builder /app/wait-for.sh .

RUN chmod +x /app/wait-for.sh

COPY .env .env

EXPOSE 8000

CMD [ "/app/task" ]