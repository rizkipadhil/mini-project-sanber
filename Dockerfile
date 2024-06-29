# Gunakan image golang versi 1.20 sebagai base image
FROM golang:1.20

# Set environment variable
ENV GO111MODULE=on

# Buat direktori kerja dalam container
WORKDIR /app

# Copy go.mod dan go.sum ke dalam container
COPY go.mod go.sum ./

# Download semua dependensi Go
RUN go mod download

# Copy semua file dari proyek ke dalam container
COPY . .

# Build aplikasi Go
RUN go build -o main .

# Eksekusi aplikasi ketika container berjalan
CMD ["./main"]

# Expose port yang digunakan aplikasi
EXPOSE 8080
