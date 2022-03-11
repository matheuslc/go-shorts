package shorturl

import (
	context "context"
	url "net/url"
)

type WriteRepository interface {
	Create(ctx context.Context, url *url.URL, alias string) (string, error)
}

func (repo RedisRepository) Create(ctx context.Context, url *url.URL, alias string) (string, error) {
	_, err := repo.client.Set(ctx, alias, url.String(), 0).Result()
	if err != nil {
		return "", err
	}

	return alias, nil
}
