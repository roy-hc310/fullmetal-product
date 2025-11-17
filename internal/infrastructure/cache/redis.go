package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/roy-hc310/fullmetal-product/pkg/config"
)

type RedisInfra struct {
	Client *redis.Client
}

func NewRedisInfra() (*RedisInfra, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         config.GlobalEnv.RedisHost,
		Password:     config.GlobalEnv.RedisPass,
		DB:           0,
		DialTimeout:  time.Second * time.Duration(config.GlobalEnv.RedisTimeOut),
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
		MaxRetries:   2,
	})

	if err := client.Ping().Err(); err != nil {
		return nil, err
	}

	return &RedisInfra{
		Client: client,
	}, nil
}

func (r *RedisInfra) Shutdown() error {
	if r.Client != nil {
		err := r.Client.Close()
		return err
	}
	return nil
}

func (r *RedisInfra) Get(ctx context.Context, key string) (string, error) {
	val, err := r.Client.WithContext(ctx).Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get key '%s': %w", key, err)
	}
	return val, nil
}
func (r *RedisInfra) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	err := r.Client.WithContext(ctx).Set(key, value, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}
func (r *RedisInfra) Delete(ctx context.Context, key string) error {
	_, err := r.Client.WithContext(ctx).Del(key).Result()
	if err != nil {
		return err
	}

	return nil
}
func (r *RedisInfra) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.Client.WithContext(ctx).Exists(key).Result()
	if err != nil {
		return false, err
	}

	// result is the number of keys that exist (1 if exists, 0 if not)
	return result > 0, nil
}
