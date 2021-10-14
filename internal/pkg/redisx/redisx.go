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

// FetchValue fetches a nondescript value from Redis.
func (r *Redis) FetchValue(ctx context.Context, key string) (interface{}, error) {
	var i interface{}
	cmd := r.r.Get(ctx, key)
	if err := cmd.Scan(i); err != nil {
		return nil, err
	}

	return i, nil
}

// SetValueWithExpiration sets a nondescript value with an expiration.
func (r *Redis) SetValueWithExpiration(ctx context.Context, key string, val interface{}, expir time.Duration) error {
	cmd := r.r.Set(ctx, key, val, expir)
	return cmd.Err()
}

// SetValue sets a nondescript value with no expiration.
func (r *Redis) SetValue(ctx context.Context, key string, val interface{}) error {
	return r.SetValueWithExpiration(ctx, key, val, 0)
}
