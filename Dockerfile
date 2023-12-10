# Gunakan image golang versi 1.8 sebagai base image
FROM golang:1.18

# Set working directory di dalam container
WORKDIR /app

# Salin file-file yang diperlukan ke dalam container
COPY XYZ-MultiFinance .

# Install dependensi jika ada
#RUN go get -d -v ./...
RUN go mod download

# Expose port 61001
EXPOSE 61001

# Command untuk menjalankan aplikasi
CMD ["go", "run", "main.go"]

