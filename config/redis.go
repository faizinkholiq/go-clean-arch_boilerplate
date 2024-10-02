package config

// import (
// 	"context"

// 	"github.com/faizinkholiq/go-clean-arch_boilerplate/config"

// 	"github.com/go-redis/redis/v8"
// )

// func InitRedis() *redis.Client {
// 	config.LoadConfig()

// 	client := redis.NewClient(&redis.Options{
// 		Addr: config.GetEnv("REDIS_HOST", "localhost:6379"),
// 	})

// 	if err := client.Ping(context.Background()).Err(); err != nil {
// 		panic("Failed to connect to Redis")
// 	}
// 	return client
// }
