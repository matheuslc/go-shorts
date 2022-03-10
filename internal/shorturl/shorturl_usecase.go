package shorturl

import (
	"context"
	"math/big"
	"net/url"
)

type ShortIntent struct {
	url *url.URL
}

type Shortener interface {
	Run(intent ShortIntent) (string, error)
}

type ShortenerService struct {
	CounterRepository
	WriteRepository
}

func (service ShortenerService) Run(intent ShortIntent) (string, error) {
	ctx := context.Background()
	nextCounter, err := service.CounterRepository.NextPosition(ctx)
	if err != nil {
		return "", nil
	}

	save, err := service.WriteRepository.Create(ctx, intent.url, base62(nextCounter))
	if err != nil {
		return "", nil
	}

	return save, nil
}

func base62(counter int64) string {
	return big.NewInt(counter).Text(62)
}
