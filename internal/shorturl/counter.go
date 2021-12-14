package shorturl

import "github.com/go-redis/redis/v8"

type CounterRepository interface {
	Next() (int32, error)
}

type RedisCounterRepository struct {
	client *redis.Client
}

type Counter struct {
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
