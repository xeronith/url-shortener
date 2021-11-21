package logging

import "xeronith/url-shortener/contracts/logging"

func CreateLogger(componentType logging.LoggerType) logging.Logger {
	switch componentType {
	case logging.DefaultLogger:
		return NewDefaultLogger()
	default:
		panic("unknown_logger_type")
	}
}
