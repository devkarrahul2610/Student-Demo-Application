# -----------------------
# Stage 1: Build Go binary
# -----------------------
FROM golang:1.25 AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o student-api ./cmd/api

# -----------------------
# Stage 2: Runtime image
# -----------------------
FROM alpine:3.18

WORKDIR /app

# Install bash (optional, nc comes from busybox)
RUN apk add --no-cache bash

# Copy binary and .env
COPY --from=builder /app/student-api .
COPY .env .

EXPOSE 8080

# Wait-for-DB before starting API
CMD sh -c "\
  until nc -z db 3306; do \
    echo 'Waiting for MySQL...'; \
    sleep 2; \
  done; \
  ./student-api \
"
