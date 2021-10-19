package redisx

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// Redis wraps *redis.Client for convenience.
type Redis struct {
	r *redis.Client
}

// Open creates a new Redis client.
func Open(addr string) *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &Redis{
		r: rdb,
	}
}

// Ping pings a Redis server.
func (r *Redis) Ping(ctx context.Context) error {
	stat := r.r.Ping(ctx)
	return stat.Err()
}

// SetStringWithExpiration sets val under key with the given expiration.
func (r *Redis) SetStringWithExpiration(ctx context.Context, key, val string, expir time.Duration) error {
	cmd := r.r.Set(ctx, key, val, expir)
	return cmd.Err()
}

// SetString sets val under key with no expiration.
func (r *Redis) SetString(ctx context.Context, key, val string) error {
	return r.SetStringWithExpiration(ctx, key, val, 0)
}

// FetchString retrieves a string from Redis.
func (r *Redis) FetchString(ctx context.Context, key string) (string, error) {
	cmd := r.r.Get(ctx, key)
	return cmd.Result()
}
