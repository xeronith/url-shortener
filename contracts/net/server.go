package net

import (
	"xeronith/url-shortener/contracts/logging"
	"xeronith/url-shortener/contracts/net/url"
)

type UrlShortenerServerType int64

const (
	DefaultUrlShortenerServer UrlShortenerServerType = iota
)

type UrlShortenerServer interface {
	Run(address string, shortener url.Shortener, logger logging.Logger) error
}
