package volatile

import (
	"errors"

	"xeronith/url-shortener/contracts/caching"
	"xeronith/url-shortener/contracts/persistence"
	_ "github.com/mattn/go-sqlite3"
)

type stringKeyValueStorage struct {
	cache caching.StringKeyValueCache
}

func NewStringKeyValueStorage(cache caching.StringKeyValueCache) persistence.StringKeyValueStorage {
	return &stringKeyValueStorage{
		cache: cache,
	}
}

func (storage *stringKeyValueStorage) Initialize() error {
	return nil
}

func (storage *stringKeyValueStorage) Persist(key, value string) error {
	storage.cache.Put(key, value)
	return nil
}

func (storage *stringKeyValueStorage) Retrieve(key string, bypassCache bool) (string, error) {
	if value, exists := storage.cache.Get(key); exists {
		return value, nil
	} else {
		return "", errors.New("key_not_found")
	}
}

func (storage *stringKeyValueStorage) Destroy() {
}
