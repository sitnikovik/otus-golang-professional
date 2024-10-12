package logger

import (
	"log/slog"
)

var (
	level Level = InfoLevel
)

// Initialize initializes the logger with the provided level
func Initialize(l Level) {
	level = l
}

// Debug logs the debug message
func Debug(msg string) {
	if level >= DebugLevel {
		slog.Debug(msg)
	}
}

// Info logs the info message
func Info(msg string) {
	if level >= InfoLevel {
		slog.Info(msg)
	}
}

// Warning logs the warning message
func Warning(msg string) {
	if level >= WarningLevel {
		slog.Warn(msg)
	}
}

// Error logs the error message
func Error(msg string) {
	if level >= ErrorLevel {
		slog.Error(msg)
	}
}

// Alert logs the alert message
func Alert(msg string) {
	if level >= AlertLevel {
		slog.Error(msg) // slog.Alert is not available
	}
}

// Critical logs the critical message
func Critical(msg string) {
	if level >= CriticalLevel {
		slog.Error(msg) // slog.Critical is not available
	}
}
