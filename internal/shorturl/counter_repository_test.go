package shorturl

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
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

		result := m.Run()
		os.Exit(result)

		return db.Ping(context.Background()).Err()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func TestNextPositionConcurrent(t *testing.T) {
	repo := RedisRepository{
		client: db,
	}

	var wg sync.WaitGroup
	var first int64
	var second int64
	var third int64

	wg.Add(3)

	go func() {
		first, _ = repo.NextPosition(context.Background())
		wg.Done()
	}()

	go func() {
		second, _ = repo.NextPosition(context.Background())
		wg.Done()
	}()

	go func() {
		third, _ = repo.NextPosition(context.Background())
		wg.Done()
	}()

	wg.Wait()

	if first == second || first == third {
		t.Errorf("Orders of counter wasnt the expect. Expected: 1, 2 and 3. Got: %d, %d and %d", &first, &second, &third)
	}
}
