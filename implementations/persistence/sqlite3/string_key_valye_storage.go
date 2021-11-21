package sqlite3

import (
	"database/sql"
	"errors"
	"os"

	"xeronith/url-shortener/contracts/caching"
	"xeronith/url-shortener/contracts/persistence"
	_ "github.com/mattn/go-sqlite3"
)

type stringKeyValueStorage struct {
	path  string
	cache caching.StringKeyValueCache
}

func NewStringKeyValueStorage(cache caching.StringKeyValueCache) persistence.StringKeyValueStorage {
	storage := &stringKeyValueStorage{
		path:  "./storage.db",
		cache: cache,
	}

	if err := storage.Initialize(); err != nil {
		panic(err)
	}

	return storage
}

func (storage *stringKeyValueStorage) Initialize() error {
	if storage.cache == nil {
		panic("missing_cache")
	}

	db, err := sql.Open("sqlite3", storage.path)
	if err != nil {
		return err
	}

	//noinspection GoUnhandledErrorResult
	defer db.Close()

	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS `string_key_values` (`key` VARCHAR(16) PRIMARY KEY, `value` VARCHAR(2048) NOT NULL);"); err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM `string_key_values`;")
	if err != nil {
		return err
	}

	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	items := make(map[string]string)
	for rows.Next() {
		var retrievedKey, retrievedValue string
		if err := rows.Scan(&retrievedKey, &retrievedValue); err != nil {
			return err
		}

		items[retrievedKey] = retrievedValue
	}

	storage.cache.Load(items)

	return nil
}

func (storage *stringKeyValueStorage) Persist(key, value string) error {
	db, err := sql.Open("sqlite3", storage.path)
	if err != nil {
		return err
	}

	//noinspection GoUnhandledErrorResult
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO `string_key_values` (`key`, `value`) VALUES (?, ?);")
	if err != nil {
		return err
	}

	//noinspection GoUnhandledErrorResult
	defer stmt.Close()

	result, err := stmt.Exec(key, value)
	if err != nil {
		return err
	}

	if affectedRows, err := result.RowsAffected(); err != nil {
		return err
	} else if affectedRows != 1 {
		return errors.New("insert_failed")
	}

	storage.cache.Put(key, value)

	return nil
}

func (storage *stringKeyValueStorage) Retrieve(key string, bypassCache bool) (string, error) {
	if !bypassCache {
		if value, exists := storage.cache.Get(key); exists {
			return value, nil
		} else {
			return "", errors.New("key_not_found")
		}
	}

	db, err := sql.Open("sqlite3", storage.path)
	if err != nil {
		return "", err
	}

	//noinspection GoUnhandledErrorResult
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `string_key_values` WHERE `key` IS ?;", key)
	if err != nil {
		return "", err
	}

	//noinspection GoUnhandledErrorResult
	defer rows.Close()

	if !rows.Next() {
		return "", errors.New("key_not_found")
	}

	var retrievedKey, retrievedValue string
	if err := rows.Scan(&retrievedKey, &retrievedValue); err != nil {
		return "", err
	}

	if retrievedKey != key {
		panic("inconsistent_keys")
	}

	return retrievedValue, nil
}

func (storage *stringKeyValueStorage) Destroy() {
	if err := os.RemoveAll(storage.path); err != nil {
		panic(err)
	}
}
