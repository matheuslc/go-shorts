package shorturl

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/go-redis/redis/v8"
	"github.com/ory/dockertest/v3"
)

var db *redis.Client
var err error

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.Run("redis", "6.2", nil)

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		db = redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp")),
		})

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func TestNextPosition(t *testing.T) {
	repo := RedisRepository{
		client: db,
	}

	result, _ := repo.NextPosition(context.Background())

	if result != 1 {
		t.Errorf("Expected string %d, got %d", 1, result)
	}
}
