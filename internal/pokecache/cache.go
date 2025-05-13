package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	res := Cache{
		entries: make(map[string]cacheEntry),
	}
	go res.reapLoop(interval)

	return res
}

func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		_ = <-ticker.C
		for key, entry := range c.entries {
			timeSinceCreation := time.Now().Sub(entry.createdAt)
			if timeSinceCreation > interval {
				c.mu.Lock()
				delete(c.entries, key)
				c.mu.Unlock()
			}
		}
	}
}