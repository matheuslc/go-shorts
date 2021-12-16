package shorturl

import (
	"context"
	"net/url"
	"testing"
)

var base10 int64 = 100
var expectedBase62 string = "1C"

type FakeRepo struct{}

func (repo FakeRepo) Next(ctx context.Context) (int64, error) {
	return base10, nil
}

func (repo FakeRepo) Create(ctx context.Context, url *url.URL, alias string) (string, error) {
	return expectedBase62, nil
}

func TestShortUrl(t *testing.T) {
	service := ShortenerService{
		CounterRepository: FakeRepo{},
		WriteRepository:   FakeRepo{},
	}

	url, _ := url.Parse("http://foo.bar")
	intent := ShortIntent{
		url: url,
	}

	shortUrl, _ := service.Run(intent)

	if shortUrl != expectedBase62 {
		t.Errorf("Expected string %s, got %s", expectedBase62, shortUrl)
	}
}
