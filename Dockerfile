# Gunakan image golang sebagai base image
FROM golang:1.20-alpine

# Set environment variables
ENV GO111MODULE=on
ENV GOPROXY=https://proxy.golang.org,direct

# Install dependencies
RUN apk add --no-cache git

# Buat direktori kerja
WORKDIR /app

# Copy go.mod dan go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy seluruh kode ke dalam container
COPY . .

# Build aplikasi
RUN go build -o main .

# Expose port yang digunakan aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./main"]
