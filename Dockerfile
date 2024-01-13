FROM golang:1.20.1-alpine AS builder

RUN apk update && \
  apk add make curl && \
  go install goa.design/goa/v3/cmd/goa@v3 && \
  go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY . .
RUN make setup
RUN chmod +x /app/wait-for.sh

EXPOSE 8000

CMD ["make", "run"]