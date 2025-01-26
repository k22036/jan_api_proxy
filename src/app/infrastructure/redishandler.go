package infrastructure

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type RedisHandler struct {
	client *redis.Client
}

func NewRedisHandler() *RedisHandler {
	return &RedisHandler{
		client: redis.NewClient(&redis.Options{
			Addr: "redis_server:6379",
		}),
	}
}

func (r *RedisHandler) Set(ctx context.Context, key string, value interface{}) error {
	err := r.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set key: %w", err)
	}
	return nil
}

func (r *RedisHandler) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get key: %w", err)
	}
	return val, nil
}
