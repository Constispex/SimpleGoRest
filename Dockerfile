# Stage 1: Build
FROM golang:1.24 AS builder
WORKDIR /app
# Abhängigkeiten herunterladen
COPY go.mod go.sum ./
RUN go mod download
# Quellcode kopieren und builden
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

# Stage 2: Minimaler Laufzeit-Room
FROM alpine:3.21.3
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Kopiere die erstellte Binary vom Builder
COPY --from=builder /app/server .
# Kopiere die .env-Datei (falls benötigt im Root)
COPY .env .

EXPOSE 8080
CMD ["./server"]
