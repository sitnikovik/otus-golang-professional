package logger

import (
	"fmt"
	"log/slog"
)

var (
	level Level = InfoLevel
)

// Initialize initializes the logger with the provided level
func Initialize(l Level) {
	level = l
}

// Debugf logs the debug with formatted message
func Debugf(format string, a ...interface{}) {
	Debug(fmt.Sprintf(format, a...))
}

// Debug logs the debug message
func Debug(msg string) {
	if level >= DebugLevel {
		slog.Debug(msg)
	}
}

// Infof logs the info with formatted message
func Infof(format string, a ...interface{}) {
	Info(fmt.Sprintf(format, a...))
}

// Info logs the info message
func Info(msg string) {
	if level >= InfoLevel {
		slog.Info(msg)
	}
}

// Warningf logs the warning with formatted message
func Warningf(format string, a ...interface{}) {
	Warning(fmt.Sprintf(format, a...))
}

// Warning logs the warning message
func Warning(msg string) {
	if level >= WarningLevel {
		slog.Warn(msg)
	}
}

// Errorf logs the error with formatted message
func Errorf(format string, a ...interface{}) {
	Warning(fmt.Sprintf(format, a...))
}

// Error logs the error message
func Error(msg string) {
	if level >= ErrorLevel {
		slog.Error(msg)
	}
}

// Alertf logs the alert with formatted message
func Alertf(format string, a ...interface{}) {
	Alert(fmt.Sprintf(format, a...))
}

// Alert logs the alert message
func Alert(msg string) {
	if level >= AlertLevel {
		slog.Error(msg) // slog.Alert is not available
	}
}

// Critical logs the critical with formatted message
func Criticalf(format string, a ...interface{}) {
	Critical(fmt.Sprintf(format, a...))
}

// Critical logs the critical message
func Critical(msg string) {
	if level >= CriticalLevel {
		slog.Error(msg) // slog.Critical is not available
	}
}
