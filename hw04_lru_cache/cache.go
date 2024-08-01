package hw04lrucache

import (
	"sync"
)

// Key is a type for cache keys.
type Key string

// Cache is an interface for a simple LRU cache.
type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

// lruCache is a simple LRU cache.
type lruCache struct {
	capacity int
	queue    List // List to store keys in LRU order
	items    map[Key]*ListItem

	mux sync.RWMutex
}

// NewCache creates a new cache with the specified capacity.
func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

// Set adds a new element with the specified value to the cache.
func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mux.Lock()
	defer c.mux.Unlock()

	if item, ok := c.items[key]; ok {
		item.Value = value
		c.queue.MoveToFront(item)
		return true
	}

	if c.queue.Len() == c.capacity {
		item := c.queue.Back()
		c.queue.Remove(item)
		delete(c.items, item.Key)
	}

	item := c.queue.PushFront(value)
	item.Key = key
	c.items[key] = item

	return false
}

// Get returns the value of the element with the specified key from the cache.
func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	if item, ok := c.items[key]; ok {
		c.queue.MoveToFront(item)
		return item.Value, true
	}

	return nil, false
}

// Clear removes all elements from the cache.
func (c *lruCache) Clear() {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
