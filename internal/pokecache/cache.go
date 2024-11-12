package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mutex    sync.Mutex
	entries  map[string]cacheEntry
	ticker   *time.Ticker
	done     chan bool
	interval time.Duration
	stopOnce sync.Once
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		entries:  make(map[string]cacheEntry),
		ticker:   time.NewTicker(interval),
		interval: interval,
		done:     make(chan bool),
	}
	go c.reapLoop()
	return c
}

func (c *Cache) reapLoop() {
	for {
		select {
		case <-c.done:
			return
		case t := <-c.ticker.C:
			c.mutex.Lock()
			for key, entry := range c.entries {
				if t.Sub(entry.createdAt) > c.interval {
					delete(c.entries, key)
				}
			}
			c.mutex.Unlock()
		}
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.entries[key]
	if !ok || time.Since(entry.createdAt) > c.interval {
		if ok {
			delete(c.entries, key)
		}
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) Shutdown(key string) {
	c.stopOnce.Do(func() {
		c.ticker.Stop()
		close(c.done)
	})
}
