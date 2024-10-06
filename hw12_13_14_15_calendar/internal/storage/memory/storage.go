package memorystorage

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

// Storage describes
type Storage interface {
	// Get returns the value by key
	Get(key string) (string, bool, error)
}

type storage struct {
	mu sync.RWMutex //nolint:unused

	redisClient *redis.Client
}

// New creates and returns the in-memory storage instance
func New(redis *redis.Client) Storage {
	return &storage{
		redisClient: redis,
	}
}

// Get returns the value by key
func (s *storage) Get(key string) (string, bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cmd := s.redisClient.Get(key)
	if cmd == nil {
		return "", false, fmt.Errorf("redis cmd is nil")
	}

	v, err := cmd.Result()
	if errors.Is(err, redis.Nil) {
		return "", false, nil
	}

	return v, true, nil
}
