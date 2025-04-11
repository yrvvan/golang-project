# 1. Build Stage
FROM golang:1.24-alpine AS builder

# Install git for module dependencies and build tools
RUN apk add --no-cache git

WORKDIR /app

# Copy the rest of the source code
COPY .env /app/.env

# Copy go.mod and go.sum to download deps early
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o server .

# 2. Run Stage (lighter image)
FROM alpine:latest

# Install CA certificates for HTTPS requests
RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy the binary from the build stage
COPY --from=builder /app/server .

# Expose your app port (adjust if your Gin app runs on another port)
EXPOSE 8089

# Run the Go server
CMD ["./server", "air"]
