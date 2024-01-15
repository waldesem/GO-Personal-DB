package cache

import (
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

// RedisConnection func for connect to Redis server.
func RedisConnection() *redis.Client {
	redisDbNum, _ := os.LookupEnv("REDIS_DB_NUMBER")
	dbNumber, _ := strconv.Atoi(redisDbNum)
	redisHost, _ := os.LookupEnv("REDIS_HOST")
	redisPort, _ := os.LookupEnv("REDIS_PORT")
	redisPswd, _ := os.LookupEnv("REDIS_PASSWORD")

	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: redisPswd,
		DB:       dbNumber,
	}

	return redis.NewClient(options)
}
