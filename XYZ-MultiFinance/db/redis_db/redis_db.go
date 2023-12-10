package redis_db

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func SetupRedisConnection() *redis.Client {
	// Membuat koneksi ke Redis
	connection := os.Getenv("REDIS_BROKER_URL")
	client := redis.NewClient(&redis.Options{
		Addr:     connection,
		Password: "",
		DB:       0,
	})

	// Menguji koneksi ke Redis dengan Context
	ctx := context.Background()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}

	return client
}
