package infrastructure

import (
	"context"
	"fmt"

	redis "github.com/redis/go-redis/v9"
)

type Redis struct {
	RDB *redis.Client
}

var ctx = context.Background()

func NewRedis() *Redis {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return &Redis{
		RDB: rdb,
	}
}

func (r *Redis) Set(key string, value interface{}) error {
	err := r.Set(key, value)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Get(key string) (value interface{}, err error) {
	value, err = r.Get(key)
	if err == redis.Nil {
		return value, fmt.Errorf("key does not exist: %v", key)
	}
	if err != nil {
		return value, err
	}
	return value, err
}
