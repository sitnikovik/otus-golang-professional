package logger

import (
	"fmt"
	"time"
)

var (
	level Level = InfoLevel
)

// SetLevel initializes the logger with the provided level
func SetLevel(l Level) {
	level = l
}

// Debugf logs the debug with formatted message
func Debugf(format string, a ...interface{}) {
	Debug(fmt.Sprintf(format, a...))
}

// Debug logs the debug message
func Debug(msg string) {
	if level <= DebugLevel {
		print(msg)
	}
}

// Infof logs the info with formatted message
func Infof(format string, a ...interface{}) {
	Info(fmt.Sprintf(format, a...))
}

// Info logs the info message
func Info(msg string) {
	if level <= InfoLevel {
		print(msg)
	}
}

// Warningf logs the warning with formatted message
func Warningf(format string, a ...interface{}) {
	Warning(fmt.Sprintf(format, a...))
}

// Warning logs the warning message
func Warning(msg string) {
	if level <= WarningLevel {
		print(msg)
	}
}

// Errorf logs the error with formatted message
func Errorf(format string, a ...interface{}) {
	Warning(fmt.Sprintf(format, a...))
}

// Error logs the error message
func Error(msg string) {
	if level <= ErrorLevel {
		print(msg)
	}
}

// Alertf logs the alert with formatted message
func Alertf(format string, a ...interface{}) {
	Alert(fmt.Sprintf(format, a...))
}

// Alert logs the alert message
func Alert(msg string) {
	if level <= AlertLevel {
		print(msg)
	}
}

// Critical logs the critical with formatted message
func Criticalf(format string, a ...interface{}) {
	Critical(fmt.Sprintf(format, a...))
}

// Critical logs the critical message
func Critical(msg string) {
	if level <= CriticalLevel {
		print(msg)
	}
}

// print logs the message with the current time
func print(msg string) {
	// no using slog package cause of diffuculties with level setting
	fmt.Printf("%s [%s]: %s\n", currTime(), level, msg)
}

// currTime returns the current time in the format "Day Mon 02 15:04:05 2006"
func currTime() string {
	return time.Now().Format(time.RFC1123)
}
