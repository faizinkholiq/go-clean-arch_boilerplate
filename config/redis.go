package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func InitRedis(cfg Config) (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return client, err
	}

	return client, nil
}
