package shorturl

import (
	context "context"
	url "net/url"
	"testing"
)

func TestCreate(t *testing.T) {
	repo := RedisRepository{
		client: db,
	}

	alias := "1C"
	urlToCreate, _ := url.Parse("http://foo.bar")

	result, err := repo.Create(context.Background(), urlToCreate, alias)
	if err != nil {
		t.Errorf(err.Error())
	}

	if result != alias {
		t.Errorf("It wasn't possible to create a url alias")
	}
}
