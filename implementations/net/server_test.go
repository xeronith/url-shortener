package net

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	. "xeronith/url-shortener/contracts/caching"
	. "xeronith/url-shortener/contracts/logging"
	. "xeronith/url-shortener/contracts/net/url"
	. "xeronith/url-shortener/contracts/persistence"
	"xeronith/url-shortener/implementations/caching"
	"xeronith/url-shortener/implementations/logging"
	"xeronith/url-shortener/implementations/net/url"
	"xeronith/url-shortener/implementations/persistence"
)

var (
	address   string
	shortener Shortener
	logger    Logger
)

func TestMain(m *testing.M) {
	address = ":8081"

	logger = logging.CreateLogger(DefaultLogger)
	cache := caching.CreateStringKeyValueCache(InMemoryStringKeyValueCache)
	storage := persistence.CreateStringKeyValueStorage(Sqlite3StringKeyValueStorage, cache)
	shortener = url.CreateShortener(DefaultShortener, "http://localhost:8081", storage)

	exitCode := m.Run()
	storage.Destroy()
	os.Exit(exitCode)
}

func Test_UrlShortenerServer_Serve(t *testing.T) {
	server := NewUrlShortenerServer()
	go func() {
		_ = server.Run(address, shortener, logger)
	}()

	originalUrl := "https://httpbin.org/get"

	reader := strings.NewReader(originalUrl)
	resp, err := http.Post(fmt.Sprintf("http://localhost%s/api/create", address), "text/plain", reader)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fail()
	}
}
