package shorturl

import (
	"context"
	"math/big"
)

type Shortener interface {
	Alias() (string, error)
}

type ShortenerService struct {
	repository CounterRepository
}

func (service ShortenerService) Alias() (string, error) {
	ctx := context.Background()
	nextCounter, err := service.repository.Next(ctx)

	if err != nil {
		return "", nil
	}

	return base62(nextCounter), nil
}

func base62(counter int64) string {
	return big.NewInt(counter).Text(62)
}
