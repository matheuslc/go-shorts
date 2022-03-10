package shorturl

import (
	"context"
	"net/url"
	"strings"
)

type RedirectIntent struct {
	url *url.URL
}

type Redirecter interface {
	Run(intent ShortIntent) (string, error)
}

type RedirectService struct {
	ReadRepository
}

func (service RedirectService) Run(intent RedirectIntent) (string, error) {
	ctx := context.Background()
	result, err := service.ReadRepository.Find(ctx, strings.Trim(intent.url.Path, "/"))
	if err != nil {
		return "", nil
	}

	return result, nil
}
