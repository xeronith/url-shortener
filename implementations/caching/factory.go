package caching

import (
	"xeronith/url-shortener/contracts/caching"
	"xeronith/url-shortener/implementations/caching/memory"
)

func CreateStringKeyValueCache(componentType caching.StringKeyValueCacheType) caching.StringKeyValueCache {
	switch componentType {
	case caching.InMemoryStringKeyValueCache:
		return memory.NewStringKeyValueCache()
	default:
		panic("unknown_string_key_value_cache_type")
	}
}
