#!/bin/bash

# Menghentikan dan menghapus semua container kecuali container bernama 'mysql'
docker-compose stop $(docker-compose ps -q | grep -v 'mysql' | awk '{print $1}')
docker-compose rm -f $(docker-compose ps -q | grep -v 'mysql' | awk '{print $1}')

# Menjalankan kembali Docker Compose
docker-compose up -d --build

