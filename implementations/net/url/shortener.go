package url

import (
	"errors"
	"fmt"
	urlParser "net/url"
	"strings"

	"xeronith/url-shortener/contracts/net/url"
	. "xeronith/url-shortener/contracts/persistence"
	stringsUtil "xeronith/url-shortener/utility/strings"
)

const (
	ShortenedLength = 6
)

type shortener struct {
	baseUrl         string
	storageProvider StringKeyValueStorage
}

func NewShortener(baseUrl string, storage StringKeyValueStorage) url.Shortener {
	if !strings.HasSuffix(baseUrl, "/") {
		baseUrl += "/"
	}

	shortener := &shortener{
		baseUrl:         baseUrl,
		storageProvider: storage,
	}

	if !shortener.UrlIsValid(shortener.baseUrl) {
		panic("invalid_base_url")
	}

	return shortener
}

func (shortener *shortener) BaseUrl() string {
	return shortener.baseUrl
}

func (shortener *shortener) UrlIsValid(url string) bool {
	if strings.TrimSpace(url) == "" {
		return false
	}

	if _, err := urlParser.ParseRequestURI(url); err != nil {
		return false
	}

	return true
}

func (shortener *shortener) Shorten(url string) (string, error) {
	if !shortener.UrlIsValid(url) {
		return "", errors.New("invalid_url")
	}

	key := stringsUtil.GenerateRandomString(ShortenedLength)
	if err := shortener.storageProvider.Persist(key, url); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s", shortener.baseUrl, key), nil
}

func (shortener *shortener) Expand(key string) (string, error) {
	if value, err := shortener.storageProvider.Retrieve(key, false); err != nil {
		return "", err
	} else {
		return value, nil
	}
}
