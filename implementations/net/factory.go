package net

import (
	"xeronith/url-shortener/contracts/net"
)

func CreateShortenerServer(componentType net.UrlShortenerServerType) net.UrlShortenerServer {
	switch componentType {
	case net.DefaultUrlShortenerServer:
		return NewUrlShortenerServer()
	default:
		panic("unknown_url_shortener_server_type")
	}
}
