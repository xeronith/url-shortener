package logging

type LoggerType int64

const (
	DefaultLogger LoggerType = iota
)

type Logger interface {
	Info(args ...interface{})
	Fail(args ...interface{})
}
