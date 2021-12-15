package shorturl

import (
	"context"
	"testing"
)

var base10 int64 = 100
var expectedBase62 string = "1C"

type FakeRepo struct{}

func (repo FakeRepo) Next(ctx context.Context) (int64, error) {
	return base10, nil
}

func TestShortUrl(t *testing.T) {
	service := ShortenerService{
		repository: FakeRepo{},
	}

	shortUrl, _ := service.Alias()

	if shortUrl != expectedBase62 {
		t.Errorf("Expected string %s, got %s", expectedBase62, shortUrl)
	}
}
