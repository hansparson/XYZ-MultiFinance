#!/bin/bash

# Menghentikan dan menghapus semua container yang sedang berjalan
docker stop $(docker ps -aq)
docker rm $(docker ps -aq)

# Menjalankan kembali Docker Compose
docker-compose up -d

# Setup prefiledge mysql
docker exec -it mysql mysql -u root -pPassword123 -e "GRANT ALL PRIVILEGES ON *.* TO 'xyz'@'mysql' IDENTIFIED BY 'Password123' WITH GRANT OPTION;"
docker exec -it mysql mysql -u root -pPassword123 -e "FLUSH PRIVILEGES;"
