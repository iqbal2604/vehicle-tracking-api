# ---------- STAGE 1: Build ----------
FROM golang:1.25.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# ---------- STAGE 2: Run ----------
FROM alpine:latest

# Buat user baru agar tidak jalan sebagai root (Security Best Practice)
RUN adduser -D appuser
WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .

# Ganti kepemilikan file ke appuser
RUN chown appuser:appuser /app/main

# Pindah ke user tersebut
USER appuser

EXPOSE 3000
CMD ["./main"]