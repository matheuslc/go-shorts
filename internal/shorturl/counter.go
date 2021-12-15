package shorturl

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type CounterRepository interface {
	Next(ctx context.Context) (int64, error)
}

type RedisCounterRepository struct {
	client *redis.Client
}

func NewRedisRepository() (RedisCounterRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})

	return RedisCounterRepository{
		client,
	}, nil
}

func (repo RedisCounterRepository) Next(ctx context.Context) (int64, error) {
	next, err := repo.client.Incr(ctx, "counter").Result()

	if err != nil {
		return -1, err
	}

	return next, nil
}
