FROM golang:1.22-alpine

# Wajib: Install build-base untuk menyediakan gcc
RUN apk add --no-cache build-base

WORKDIR /app

# Wajib: Aktifkan CGO agar flag -race bisa bekerja
ENV CGO_ENABLED=1

COPY . .
RUN go mod tidy

# Build aplikasi sesuai struktur folder Anda [cite: 32]
RUN go build -o app ./cmd/server

CMD ["./app"]