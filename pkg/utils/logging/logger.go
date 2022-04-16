package logging

import (
	"log"
	"os"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}

type BuiltinLogger struct {
	logger *log.Logger
}

func NewBuiltinLogger() *BuiltinLogger {
	return &BuiltinLogger{logger: log.New(os.Stdout, "", 5)}
}

func (l *BuiltinLogger) Debug(args ...interface{}) {
	l.logger.Println(args...)
}

func (l *BuiltinLogger) Debugf(format string, args ...interface{}) {
	l.logger.Printf(format, args...)
}
