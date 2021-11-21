package url

import (
	"xeronith/url-shortener/contracts/net/url"
	"xeronith/url-shortener/contracts/persistence"
)

func CreateShortener(componentType url.ShortenerType, baseUrl string, storage persistence.StringKeyValueStorage) url.Shortener {
	switch componentType {
	case url.DefaultShortener:
		return NewShortener(baseUrl, storage)
	default:
		panic("unknown_shortener_type")
	}
}
