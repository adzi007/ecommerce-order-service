# Builder stage
FROM golang:1.23.4-alpine AS builder

# Install git and required tools for Go modules
RUN apk add --no-cache git

WORKDIR /app

# Copy go mod files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the main application
RUN go build -o main ./cmd/main.go

# Build the migration binary
RUN go build -o migrate ./internal/migration/migration.go

# Final stage (minimal image)
FROM alpine:latest

# Install SSL certificates (needed for HTTPS or DB connections)
RUN apk add --no-cache ca-certificates

# Copy compiled binaries and env
COPY --from=builder /app/main /main
COPY --from=builder /app/migrate /migrate
COPY --from=builder /app/.env /app/.env

# Set working directory
WORKDIR /app

# Set default command
CMD ["/main"]
