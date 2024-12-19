package logger

import (
	"strings"
)

// Level defines the log level.
type Level int

// level is the current logging level.
var level = InfoLevel

// SetLevel initializes the logger with the provided level.
func SetLevel(l Level) {
	level = l
}

const (
	// EmergencyLevel level is used for emergency messages.
	EmergencyLevel Level = iota
	// AlertLevel level is used for alerts.
	AlertLevel
	// CriticalLevel level is used for critical messages.
	CriticalLevel
	// ErrorLevel level is used for errors.
	ErrorLevel
	// WarningLevel level is used for warnings.
	WarningLevel
	// NoticeLevel level is used for warnings.
	NoticeLevel
	// InfoLevel level is used for informational messages.
	InfoLevel
	// DebugLevel level is used for debugging purposes.
	DebugLevel
)

// String returns the string representation of the log level.
func (l Level) String() string {
	switch l {
	case EmergencyLevel:
		return "emergency"
	case AlertLevel:
		return "alert"
	case CriticalLevel:
		return "critical"
	case ErrorLevel:
		return "error"
	case WarningLevel:
		return "warning"
	case NoticeLevel:
		return "notice"
	case InfoLevel:
		return "info"
	case DebugLevel:
		return "debug"
	}

	return ""
}

// LevelFromString returns the log level from the string.
func LevelFromString(level string) Level {
	switch strings.ToLower(level) {
	case "emergency":
		return EmergencyLevel
	case "alert":
		return AlertLevel
	case "critical":
		return CriticalLevel
	case "error":
		return ErrorLevel
	case "warning":
		return WarningLevel
	case "notice":
		return NoticeLevel
	case "info":
		return InfoLevel
	case "debug":
		return DebugLevel
	}

	return InfoLevel
}
