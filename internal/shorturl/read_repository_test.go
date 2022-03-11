package shorturl

import (
	"context"
	url "net/url"
	"testing"
)

func TestFindAgain(t *testing.T) {
	repo := RedisRepository{
		client: db,
	}

	alias := "1C"
	urlToCreate, _ := url.Parse("http://foo.bar")

	// Creating
	repo.Create(context.Background(), urlToCreate, alias)

	found, _ := repo.Find(context.Background(), alias)

	if found != urlToCreate.String() {
		t.Errorf("It wasn't possible to create a url alias: %s", found)
	}
}
