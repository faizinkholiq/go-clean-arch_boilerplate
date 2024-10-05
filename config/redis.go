package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func InitRedis() (client *redis.Client, err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     GetConf.Redis.Host,
		Password: GetConf.Redis.Password,
		DB:       GetConf.Redis.DB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return client, err
	}

	return client, nil
}
