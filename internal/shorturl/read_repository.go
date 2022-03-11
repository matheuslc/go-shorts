package shorturl

import context "context"

type ReadRepository interface {
	Find(ctx context.Context, alias string) (string, error)
}

func (repo RedisRepository) Find(ctx context.Context, alias string) (string, error) {
	result, err := repo.client.Get(ctx, alias).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}
