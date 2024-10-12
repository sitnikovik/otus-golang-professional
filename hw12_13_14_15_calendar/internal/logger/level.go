package logger

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

func LevelFromString(level string) Level {
	switch level {
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
