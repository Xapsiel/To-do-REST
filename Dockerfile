# Многоступенчатая сборка
FROM golang:1.22.3 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy

COPY . ./
RUN go build -o /main ./cmd/main.go


EXPOSE 8080
ENTRYPOINT ["/main"]