package shorturl

import (
	"context"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
)

var base10 int64 = 100
var expectedBase62 string = "1C"

func TestShortUrl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	url, _ := url.Parse("http://foo.bar")

	fakerCounterRepo := NewMockCounterRepository(ctrl)
	fakerWriterRepo := NewMockWriteRepository(ctrl)

	fakerCounterRepo.
		EXPECT().
		NextPosition(context.Background()).
		Return(base10, nil)

	fakerWriterRepo.
		EXPECT().
		Create(context.Background(), gomock.Eq(url), expectedBase62).
		Return(expectedBase62, nil)

	service := ShortenerService{
		CounterRepository: fakerCounterRepo,
		WriteRepository:   fakerWriterRepo,
	}

	intent := ShortIntent{
		url: url,
	}

	shortUrl, _ := service.Run(intent)

	if shortUrl != expectedBase62 {
		t.Errorf("Expected string %s, got %s", expectedBase62, shortUrl)
	}
}
