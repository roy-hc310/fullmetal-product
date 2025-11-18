# ---- BUILD STAGE ----
FROM golang:1.25.0-alpine AS builder

WORKDIR /app

# Install build deps (optional but recommended)
# RUN apk add --no-cache git

# Copy go.mod and go.sum first for better layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the full project
COPY . .

# Build the binary (note path to main.go)
RUN CGO_ENABLED=0 GOOS=linux go build -o fullmetal-product ./cmd/server

# ---- RUNTIME STAGE ----
FROM alpine:latest

WORKDIR /root

# Add certificates (important for HTTP(S)/gRPC calls)
# RUN apk add --no-cache ca-certificates

# Copy binary from builder
COPY --from=builder /app/fullmetal-product .

# Expose gRPC port
EXPOSE 50051

# Command to run
CMD ["./fullmetal-product"]