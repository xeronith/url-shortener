package net

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"xeronith/url-shortener/contracts/logging"
	"xeronith/url-shortener/contracts/net"
	. "xeronith/url-shortener/contracts/net/url"
)

type urlShortenerServer struct {
	shortener Shortener
	logger    logging.Logger
}

func NewUrlShortenerServer() net.UrlShortenerServer {
	return &urlShortenerServer{}
}

func (server *urlShortenerServer) Run(address string, shortener Shortener, logger logging.Logger) error {
	if shortener == nil {
		panic("shortener_required")
	}

	if logger == nil {
		panic("logger_required")
	}

	server.shortener = shortener
	server.logger = logger

	http.HandleFunc("/api/create", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "POST":
			body, err := ioutil.ReadAll(request.Body)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}

			shortenedUrl, err := server.shortener.Shorten(string(body))
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
			}

			writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
			writer.Header().Set("X-Content-Type-Options", "nosniff")
			writer.WriteHeader(http.StatusOK)
			//noinspection GoUnhandledErrorResult
			writer.Write([]byte(shortenedUrl))
		default:
			http.Error(writer, "method_not_allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		//noinspection GoUnhandledErrorResult
		writer.Write([]byte("success"))
	})

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			shortenedUrlKey := strings.TrimPrefix(request.URL.Path, "/")
			if expandedUrl, err := server.shortener.Expand(shortenedUrlKey); err != nil {
				http.Error(writer, err.Error(), http.StatusNotFound)
			} else {
				http.Redirect(writer, request, expandedUrl, http.StatusMovedPermanently)
			}
		default:
			http.Error(writer, "method_not_allowed", http.StatusMethodNotAllowed)
		}
	})

	server.logger.Info(fmt.Sprintf("Server started on %s with %s as base url", address, server.shortener.BaseUrl()))
	return http.ListenAndServe(address, nil)
}
