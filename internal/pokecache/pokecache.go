package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	data     map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	var cache Cache
	cache.data = map[string]cacheEntry{}
	cache.interval = interval
	cache.reapLoop()
	return cache
}

func (cache Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	var newEntry = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	cache.data[key] = newEntry

}

func (cache Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	val, ok := cache.data[key]
	if !ok {
		return nil, false
	}
	return val.val, true
}

func (cache Cache) reapLoop() {
	ticker := time.NewTicker(cache.interval)

	go func() {
		for {
			<-ticker.C
			cache.mu.Lock()
			for key, ce := range cache.data {
				if time.Now().Sub(ce.createdAt) > cache.interval {
					delete(cache.data, key)
				}
			}
			cache.mu.Unlock()
		}
	}()
}
