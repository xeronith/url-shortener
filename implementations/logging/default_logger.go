package logging

import (
	"log"

	"xeronith/url-shortener/contracts/logging"
)

type defaultLogger struct{}

func NewDefaultLogger() logging.Logger {
	return &defaultLogger{}
}

func (logger *defaultLogger) Info(args ...interface{}) {
	log.Println(args...)
}

func (logger *defaultLogger) Fail(args ...interface{}) {
	log.Fatalln(args...)
}
