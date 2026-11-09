# --- STAGE 1: Build ---
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Static binary, no debug info, no local paths baked in
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /app/server .

# --- STAGE 2: Runtime ---
FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /app/server /server

EXPOSE 8080

USER nonroot:nonroot
ENTRYPOINT ["/server"]