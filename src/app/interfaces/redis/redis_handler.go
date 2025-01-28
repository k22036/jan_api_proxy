package redis

import "context"

type RedisHandler interface {
	Set(ctx context.Context, key string, value interface{}) error
	Get(ctx context.Context, key string) (string, error)
	MGet(ctx context.Context, keys ...string) ([]string, error)
	AllKeys(ctx context.Context) ([]string, error)
}
