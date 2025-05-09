package logger

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

type Logger interface {
	Debug(msg string, fields ...any)
	Info(msg string, fields ...any)
	Warn(msg string, fields ...any)
	Error(msg string, fields ...any)
	Fatal(msg string, fields ...any)
}

// Allows logging with golang-migrate
type MigrateLogger interface {
  Logger
  Printf(msg string, fields ...any)
  Verbose() bool
}
