# Gunakan image Golang versi alpine (ringan)
FROM golang:1.25.5-alpine

# Set folder kerja di dalam container
WORKDIR /app

#install air untuk hot reload
RUN go install github.com/air-verse/air@latest

# Copy file dependency
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build aplikasi
RUN go build -o main .

# Expose port yang digunakan
EXPOSE 3000

# Jalankan aplikasi dengan air untuk hot reload
CMD ["air"]