package logger

import (
	"strings"
)

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

// String returns the string representation of the log level
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarningLevel:
		return "WARNING"
	case ErrorLevel:
		return "ERROR"
	case AlertLevel:
		return "ALERT"
	case CriticalLevel:
		return "CRITICAL"
	default:
		return "UNKNOWN"
	}
}

func LevelFromString(level string) Level {
	lvl := strings.ToLower(level)
	switch lvl {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warning":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "alert":
		return AlertLevel
	case "critical":
		return CriticalLevel
	default:
		return InfoLevel
	}
}
