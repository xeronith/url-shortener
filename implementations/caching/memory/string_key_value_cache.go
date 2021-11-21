package memory

import (
	"sync"

	"xeronith/url-shortener/contracts/caching"
)

type stringKeyValueCache struct {
	mutex      sync.RWMutex
	collection map[string]string
}

func NewStringKeyValueCache() caching.StringKeyValueCache {
	return &stringKeyValueCache{
		collection: make(map[string]string),
	}
}

func (cache *stringKeyValueCache) Load(items map[string]string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.collection = items
}

func (cache *stringKeyValueCache) Put(key, value string) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()

	cache.collection[key] = value
}

func (cache *stringKeyValueCache) Get(key string) (string, bool) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()

	if value, exists := cache.collection[key]; !exists {
		return "", false
	} else {
		return value, true
	}
}
