package internal

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const TTL1Week = 7 * 24 * 60 * 60 * time.Second

type RedisCache struct {
	client *redis.Client
}

// NewRedisCache creates a Redis-backed cache using the given address (e.g. "localhost:6379").
func NewRedisCache(addr string) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{Addr: addr})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("redis ping: %w", err)
	}
	return &RedisCache{client: client}, nil
}

func (r *RedisCache) Get(ctx context.Context, key string) (bool, error) {
	_, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *RedisCache) Set(ctx context.Context, key string) error {
	return r.client.Set(ctx, key, key, TTL1Week).Err()
}

func (r *RedisCache) Close() error {
	return r.client.Close()
}
