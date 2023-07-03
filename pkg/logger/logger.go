package logger

type Logger interface {
	Warn(kv ...interface{})
	Error(kv ...interface{})
	Debug(kv ...interface{})
	Info(kv ...interface{})
}
