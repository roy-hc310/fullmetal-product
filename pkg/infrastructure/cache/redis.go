package cache

import (
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
