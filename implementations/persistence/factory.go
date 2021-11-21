package persistence

import (
	"xeronith/url-shortener/contracts/caching"
	"xeronith/url-shortener/contracts/persistence"
	"xeronith/url-shortener/implementations/persistence/sqlite3"
	"xeronith/url-shortener/implementations/persistence/volatile"
)

func CreateStringKeyValueStorage(componentType persistence.StringKeyValueStorageType, cache caching.StringKeyValueCache) persistence.StringKeyValueStorage {
	switch componentType {
	case persistence.VolatileStringKeyValueStorage:
		return volatile.NewStringKeyValueStorage(cache)
	case persistence.Sqlite3StringKeyValueStorage:
		return sqlite3.NewStringKeyValueStorage(cache)
	default:
		panic("unknown_string_key_value_storage_type")
	}
}
