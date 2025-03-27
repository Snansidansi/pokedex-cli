package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	data      []byte
}

type Cache struct {
	entries       map[string]cacheEntry
	mu            *sync.Mutex
	checkInterval time.Duration
}

func NewCache(checkInterval time.Duration, maxCacheAge time.Duration) Cache {
	cache := Cache{
		entries:       make(map[string]cacheEntry),
		mu:            &sync.Mutex{},
		checkInterval: checkInterval,
	}

	go cache.reapLoop(maxCacheAge)
	return cache
}

func (c *Cache) Add(key string, data []byte) {
	cacheEntry := cacheEntry{
		createdAt: time.Now(),
		data:      data,
	}

	c.mu.Lock()
	c.entries[key] = cacheEntry
	c.mu.Unlock()
}
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	val, ok := c.entries[key]
	c.mu.Unlock()

	if ok == false {
		return nil, false
	}
	return val.data, true
}

func (c *Cache) reapLoop(maxCacheAge time.Duration) {
	ticker := time.NewTicker(c.checkInterval)
	defer ticker.Stop()

	for {
		<-ticker.C

		c.mu.Lock()
		for key, cacheEntry := range c.entries {
			cacheAge := time.Now().Sub(cacheEntry.createdAt)
			if cacheAge > maxCacheAge {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
