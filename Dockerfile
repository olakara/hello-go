# Build stage
FROM golang:1.23-alpine AS builder

LABEL authors="olakara"
LABEL maintainer="olakara"
LABEL description="This is a simple Dockerfile for building a Go application"
LABEL version="1.0"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server .

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 3000
ENTRYPOINT ["./server"]