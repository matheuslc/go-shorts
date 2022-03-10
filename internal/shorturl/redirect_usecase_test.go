package shorturl

import (
	"context"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestRedirecService(t *testing.T) {
	var expectedRedirect string = "http://redirect.com"
	var expectedBase62 string = "1C"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	url, _ := url.Parse("http://foo.bar/1C")

	fakerCounterRepo := NewMockReadRepository(ctrl)

	fakerCounterRepo.
		EXPECT().
		Find(context.Background(), expectedBase62).
		Return(expectedRedirect, nil)

	service := RedirectService{
		ReadRepository: fakerCounterRepo,
	}

	intent := RedirectIntent{
		url: url,
	}

	shortUrl, _ := service.Run(intent)

	if shortUrl != expectedRedirect {
		t.Errorf("Expected string %s, got %s", expectedRedirect, shortUrl)
	}
}
