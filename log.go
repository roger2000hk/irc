package irc

import (
	"log"
)

// Simple logger interface designed for use with logrus
// and other similar systems
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Print(args ...interface{})
}

// AKA Black hole logger or /dev/null logger.
type NilLogger struct{}

func (l *NilLogger) Debug(args ...interface{}) {
}

func (l *NilLogger) Info(args ...interface{}) {
}

func (l *NilLogger) Warn(args ...interface{}) {
}

func (l *NilLogger) Error(args ...interface{}) {
}

func (l *NilLogger) Fatal(args ...interface{}) {
}

func (l *NilLogger) Print(args ...interface{}) {
}

// This logger simply tries to proxy to the official golang log package
type SimpleLogger struct{}

func (l *SimpleLogger) Debug(args ...interface{}) {
	data := append([]interface{}{"DEBUG"}, args...)
	log.Print(data...)
}

func (l *SimpleLogger) Info(args ...interface{}) {
	data := append([]interface{}{"INFO"}, args...)
	log.Print(data...)
}

func (l *SimpleLogger) Warn(args ...interface{}) {
	data := append([]interface{}{"WARN"}, args...)
	log.Print(data...)
}

func (l *SimpleLogger) Error(args ...interface{}) {
	data := append([]interface{}{"ERROR"}, args...)
	log.Print(data...)
}

func (l *SimpleLogger) Fatal(args ...interface{}) {
	data := append([]interface{}{"FATAL"}, args...)
	log.Fatal(data...)
}

func (l *SimpleLogger) Print(args ...interface{}) {
	log.Print(args...)
}
