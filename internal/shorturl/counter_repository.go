package shorturl

import (
	"context"

	redis "github.com/go-redis/redis/v8"
	"github.com/matheuslc/go-shorts/config"
)

type CounterRepository interface {
	NextPosition(ctx context.Context) (int64, error)
}

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(cfg config.RedisOptions) (RedisRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddress,
		Password: cfg.RedisPassword,
		DB:       cfg.Database,
	})

	return RedisRepository{
		client,
	}, nil
}

func (repo RedisRepository) NextPosition(ctx context.Context) (int64, error) {
	next, err := repo.client.Incr(ctx, "counter").Result()

	if err != nil {
		return -1, err
	}

	return next, nil
}
