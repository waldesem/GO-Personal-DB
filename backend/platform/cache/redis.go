package cache

import (
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// RedisConnection func for connect to Redis server.
func RedisConnection() *redis.Client {
	redisDbNum := os.Getenv("REDIS_DB_NUMBER")
	dbNumber, _ := strconv.Atoi(redisDbNum)
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPswd := os.Getenv("REDIS_PASSWORD")

	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPswd,
		DB:       dbNumber,
	}

	return redis.NewClient(options)
}
