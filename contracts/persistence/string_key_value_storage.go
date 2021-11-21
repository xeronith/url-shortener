package persistence

type StringKeyValueStorageType int64

const (
	VolatileStringKeyValueStorage StringKeyValueStorageType = iota
	Sqlite3StringKeyValueStorage
)

type StringKeyValueStorage interface {
	Initialize() error
	Persist(key, value string) error
	Retrieve(key string, bypassCache bool) (string, error)
	Destroy()
}
