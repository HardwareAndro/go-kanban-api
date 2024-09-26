FROM golang:1.22.7-alpine3.20 AS builder

WORKDIR /app
COPY . .

RUN go build

CMD ["./go-kanban-api"]