package shorturl

import (
	"context"
	"net/url"

	"github.com/go-redis/redis/v8"
)

type CounterRepository interface {
	Next(ctx context.Context) (int64, error)
}

type WriteRepository interface {
	Create(ctx context.Context, url *url.URL, alias string) (string, error)
}

type ReadRepository interface {
	Find(ctx context.Context, alias string) (string, error)
}

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository() (RedisRepository, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "",
		Password: "",
		DB:       0,
	})

	return RedisRepository{
		client,
	}, nil
}

func (repo RedisRepository) Next(ctx context.Context) (int64, error) {
	next, err := repo.client.Incr(ctx, "counter").Result()

	if err != nil {
		return -1, err
	}

	return next, nil
}

func (repo RedisRepository) Create(ctx context.Context, url *url.URL, alias string) (string, error) {
	_, err := repo.client.Set(ctx, alias, url.String(), 0).Result()
	if err != nil {
		return "", err
	}

	return alias, nil
}

func (repo RedisRepository) Find(ctx context.Context, alias string) (string, error) {
	result, err := repo.client.Get(ctx, alias).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}
