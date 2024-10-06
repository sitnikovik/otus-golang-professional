package config

// LoggerConf describes the logging configuration
type LoggerConf struct {
	Level string
}

// NewLoggerConf creates and returns the instance to configure app logging
func NewLoggerConf(level string) LoggerConf {
	return LoggerConf{
		Level: level,
	}
}
