package caching

type StringKeyValueCacheType int64

const (
	InMemoryStringKeyValueCache StringKeyValueCacheType = iota
)

type StringKeyValueCache interface {
	Load(map[string]string)
	Put(key, value string)
	Get(key string) (string, bool)
}
