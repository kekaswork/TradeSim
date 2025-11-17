FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git for go modules
RUN apk add --no-cache git

# Copy go mod files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o auth-service ./cmd/auth

# ============================
# 2. Final runtime stage
# ============================
FROM alpine:3.19

WORKDIR /app

# Add certificates for HTTPS requests
RUN apk add --no-cache ca-certificates

# Copy compiled binary
COPY --from=builder /app/auth-service .

# Expose port
EXPOSE 8080

CMD ["./auth-service"]