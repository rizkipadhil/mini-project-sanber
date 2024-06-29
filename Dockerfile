# Gunakan image base yang lebih kecil karena kita hanya perlu menjalankan binary
FROM golang:1.20-alpine

# Buat direktori kerja dalam container
WORKDIR /app

# Copy file hasil build ke dalam container
COPY main .

# Eksekusi aplikasi ketika container berjalan
CMD ["./main"]

# Expose port yang digunakan aplikasi
EXPOSE 8080
