# --- STAGE 1: Build the application ---
FROM golang:1.23-alpine AS builder

# Install git if your dependencies require it to pull from private repos
RUN apk add --no-cache git
WORKDIR /app

# Copy dependency manifests first to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]
