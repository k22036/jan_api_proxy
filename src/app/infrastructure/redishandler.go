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

func (r *RedisHandler) MGet(ctx context.Context, keys ...string) ([]string, error) {
	vals, err := r.client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get keys: %w", err)
	}

	var results []string
	for _, val := range vals {
		if val == nil {
			results = append(results, "")
			continue
		}
		results = append(results, val.(string))
	}
	return results, nil
}

func (r *RedisHandler) AllKeys(ctx context.Context) ([]string, error) {
	keys, err := r.client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get all keys: %w", err)
	}
	return keys, nil
}

func (r *RedisHandler) Delete(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("failed to delete key: %w", err)
	}
	return nil
}
