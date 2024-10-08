package logger

import "fmt"

// Level defines the log level
type Level int

const (
	// DebugLevel level is used for debugging purposes
	DebugLevel Level = iota + 1
	// InfoLevel level is used for informational messages
	InfoLevel
	// WarningLevel level is used for warnings
	WarningLevel
	// ErrorLevel level is used for errors
	ErrorLevel
	// AlertLevel level is used for alerts
	AlertLevel
	// CriticalLevel level is used for critical messages
	CriticalLevel
)

var (
	level Level = InfoLevel
)

// Init initializes the logger with the provided level
func Initialize(l Level) {
	level = l
}

// Debug logs the debug message
func Debug(msg string) {
	if level >= DebugLevel {
		fmt.Println(msg)
	}
}

// Info logs the info message
func Info(msg string) {
	if level >= InfoLevel {
		fmt.Println(msg)
	}
}

// Warning logs the warning message
func Warning(msg string) {
	if level >= WarningLevel {
		fmt.Println(msg)
	}
}

// Error logs the error message
func Error(msg string) {
	if level >= ErrorLevel {
		fmt.Println(msg)
	}
}

// Alert logs the alert message
func Alert(msg string) {
	if level >= AlertLevel {
		fmt.Println(msg)
	}
}

// Critical logs the critical message
func Critical(msg string) {
	if level >= CriticalLevel {
		fmt.Println(msg)
	}
}
