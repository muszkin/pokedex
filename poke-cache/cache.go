package poke_cache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache    map[string]CacheEntry
	mut      *sync.RWMutex
	interval time.Duration
}

func (c *Cache) Add(key string, value []byte) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.cache[key] = CacheEntry{val: value, createdAt: time.Now()}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mut.RLock()
	defer c.mut.RUnlock()
	cacheEntry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return cacheEntry.val, true
}

func (c *Cache) reapLoop(currentTime time.Time) {
	c.mut.Lock()
	defer c.mut.Unlock()
	for i, item := range c.cache {
		duration := currentTime.Sub(item.createdAt)
		if duration > c.interval {
			delete(c.cache, i)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	var cacheMap = make(map[string]CacheEntry)
	var mu = sync.RWMutex{}
	ticker := time.NewTicker(interval)
	c := Cache{cache: cacheMap, mut: &mu, interval: interval}
	go func() {
		for {
			select {
			case currentTime := <-ticker.C:
				c.reapLoop(currentTime)
			}
		}
	}()
	return c
}
