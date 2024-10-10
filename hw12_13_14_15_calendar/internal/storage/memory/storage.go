package memorystorage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"github.com/go-redis/redis"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/storage"
)

// Storage describes
type Storage interface {
	// CreateEvent creates a new event
	CreateEvent(ctx context.Context, event *storage.Event) error
	// UpdateEvent updates the event
	UpdateEvent(ctx context.Context, event *storage.Event) error
	// DeleteEvent deletes the event
	DeleteEvent(ctx context.Context, eventID string) error
	// GetEvent returns the event by ID
	GetEvent(ctx context.Context, eventID string) (*storage.Event, error)

	// Close closes the storage
	Close() error
}

type redisStorage struct {
	mu sync.RWMutex //nolint:unused

	redisClient *redis.Client
}

// New creates and returns the in-memory storage instance
func NewRedis(redis *redis.Client) Storage {
	return &redisStorage{
		redisClient: redis,
	}
}

// Close closes the storage
func (s *redisStorage) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.redisClient.Close()
}

// CreateEvent creates a new event
func (s *redisStorage) CreateEvent(ctx context.Context, event *storage.Event) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	return s.set(generateKey(event), string(jsonData))
}

// UpdateEvent updates the event
func (s *redisStorage) UpdateEvent(ctx context.Context, event *storage.Event) error {
	// Check if the event exists
	exists, err := s.exists(generateKey(event))
	if err != nil {
		return fmt.Errorf("failed to check if the event exists: %w", err)
	}
	if !exists {
		return storage.ErrNotFound
	}

	// Update the event
	jsonData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	return s.set(generateKey(event), string(jsonData))
}

// DeleteEvent deletes the event
func (s *redisStorage) DeleteEvent(ctx context.Context, eventID string) error {
	return s.del(generateKey(&storage.Event{ID: eventID}))
}

// GetEvent returns the event by ID
func (s *redisStorage) GetEvent(ctx context.Context, eventID string) (*storage.Event, error) {
	v, ok, err := s.get(generateKey(&storage.Event{ID: eventID}))
	if err != nil {
		return nil, fmt.Errorf("failed to get event: %w", err)
	}

	if !ok {
		return nil, storage.ErrNotFound
	}

	var event storage.Event
	if err := json.Unmarshal([]byte(v), &event); err != nil {
		return nil, fmt.Errorf("failed to unmarshal event: %w", err)
	}

	return &event, nil
}

// get returns the value by key
func (s *redisStorage) get(key string) (string, bool, error) {
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

// exists checks if the key exists
func (s *redisStorage) exists(key string) (bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cmd := s.redisClient.Exists(key)
	if cmd == nil {
		return false, fmt.Errorf("redis cmd is nil")
	}

	v, err := cmd.Result()
	if err != nil {
		return false, fmt.Errorf("failed to get key: %w", err)
	}

	return v > 0, nil
}

// set sets the value by key
func (s *redisStorage) set(key, value string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	cmd := s.redisClient.Set(key, value, 0)
	if cmd == nil {
		return fmt.Errorf("redis cmd is nil")
	}

	return nil
}

// del deletes the value by key
func (s *redisStorage) del(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	cmd := s.redisClient.Del(key)
	if cmd == nil {
		return fmt.Errorf("redis cmd is nil")
	}

	return nil
}

// generateKey generates the key for the event
func generateKey(event *storage.Event) string {
	return fmt.Sprintf("event.%s", event.ID)
}
