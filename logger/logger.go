// logger/logger.go
package logger

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Logger struct {
	level   LogLevel
	mu      sync.Mutex
	writers []io.Writer
}

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

func NewLogger(level LogLevel, writers ...io.Writer) (*Logger, error) {
	if len(writers) == 0 {
		return nil, fmt.Errorf("at least one writer is required")
	}
	return &Logger{level: level, writers: writers}, nil
}

func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if level < l.level {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] [%s] %s\n", timestamp, levelString(level), fmt.Sprintf(format, args...))

	for _, w := range l.writers {
		_, err := io.WriteString(w, logEntry)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write log entry: %v\n", err)
		}
	}
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.log(DEBUG, format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.log(INFO, format, args...)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.log(WARN, format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.log(ERROR, format, args...)
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	l.log(FATAL, format, args...)
	os.Exit(1)
}

func levelString(level LogLevel) string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}
