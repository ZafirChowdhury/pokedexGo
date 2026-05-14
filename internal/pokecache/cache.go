package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mutex *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: map[string]cacheEntry{},
		mutex: &sync.Mutex{},
	}

	go c.cleanupLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		value:     val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	val, ok := c.cache[key]
	return val.value, ok
}

func (c *Cache) cleanupLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.cleanup(time.Now().UTC(), interval)
	}
}

func (c *Cache) cleanup(now time.Time, interval time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for key, val := range c.cache {
		if val.createdAt.Before(now.Add(-interval)) {
			delete(c.cache, key)
		}
	}
}
