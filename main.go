package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	. "xeronith/url-shortener/contracts/caching"
	. "xeronith/url-shortener/contracts/logging"
	. "xeronith/url-shortener/contracts/net"
	. "xeronith/url-shortener/contracts/net/url"
	. "xeronith/url-shortener/contracts/persistence"
	"xeronith/url-shortener/implementations/caching"
	"xeronith/url-shortener/implementations/logging"
	"xeronith/url-shortener/implementations/net"
	"xeronith/url-shortener/implementations/net/url"
	"xeronith/url-shortener/implementations/persistence"
)

var (
	baseUrlFlag = flag.String("b", "http://localhost:8080", "Base url of the server including the protocol, for example 'https://yourdomain.com'")
	portFlag    = flag.Int("p", 8080, "Exposed web application port")
)

func main() {
	flag.Parse()

	if os.Getenv("BASE_URL") != "" {
		*baseUrlFlag = os.Getenv("BASE_URL")
	}

	if os.Getenv("PORT") != "" {
		if port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 32); err == nil {
			*portFlag = int(port)
		}
	}

	server := net.CreateShortenerServer(DefaultUrlShortenerServer)
	logger := logging.CreateLogger(DefaultLogger)
	cache := caching.CreateStringKeyValueCache(InMemoryStringKeyValueCache)
	storage := persistence.CreateStringKeyValueStorage(Sqlite3StringKeyValueStorage, cache)
	shortener := url.CreateShortener(DefaultShortener, *baseUrlFlag, storage)

	if err := server.Run(fmt.Sprintf(":%d", *portFlag), shortener, logger); err != nil {
		panic(err)
	}
}
