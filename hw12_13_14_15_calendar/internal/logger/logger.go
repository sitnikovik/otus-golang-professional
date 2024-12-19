package logger

import (
	"fmt"
	"time"
)

// Debugf logs the debug with formatted message
func Debugf(format string, a ...interface{}) {
	Debug(fmt.Sprintf(format, a...))
}

// Debug logs the debug message
func Debug(msg string) {
	if level >= DebugLevel {
		log(DebugLevel, msg)
	}
}

// Infof logs the info with formatted message
func Infof(format string, a ...interface{}) {
	Info(fmt.Sprintf(format, a...))
}

// Info logs the info message
func Info(msg string) {
	if level >= InfoLevel {
		log(InfoLevel, msg)
	}
}

// WarnNoticefingf logs the notice with formatted message
func Noticef(format string, a ...interface{}) {
	Notice(fmt.Sprintf(format, a...))
}

// Notice logs the notice message
func Notice(msg string) {
	if level >= NoticeLevel {
		log(NoticeLevel, msg)
	}
}

// Warningf logs the warning with formatted message
func Warningf(format string, a ...interface{}) {
	Warning(fmt.Sprintf(format, a...))
}

// Warning logs the warning message
func Warning(msg string) {
	if level >= WarningLevel {
		log(WarningLevel, msg)
	}
}

// Errorf logs the error with formatted message
func Errorf(format string, a ...interface{}) {
	Error(fmt.Sprintf(format, a...))
}

// Error logs the error message
func Error(msg string) {
	if level >= ErrorLevel {
		log(ErrorLevel, msg)
	}
}

// Alertf logs the alert with formatted message
func Alertf(format string, a ...interface{}) {
	Alert(fmt.Sprintf(format, a...))
}

// Alert logs the alert message
func Alert(msg string) {
	if level >= AlertLevel {
		log(AlertLevel, msg)
	}
}

// Critical logs the critical with formatted message
func Criticalf(format string, a ...interface{}) {
	Critical(fmt.Sprintf(format, a...))
}

// Critical logs the critical message
func Critical(msg string) {
	if level >= CriticalLevel {
		log(CriticalLevel, msg)
	}
}

// Emergency logs the emergency with formatted message
func Emergencyf(format string, a ...interface{}) {
	Emergency(fmt.Sprintf(format, a...))
}

// Emergency logs the emergency message
func Emergency(msg string) {
	if level >= EmergencyLevel {
		log(EmergencyLevel, msg)
		panic(msg)
	}
}

// log logs the message
func log(lvl Level, msg string) {
	// no using slog package cause of diffuculties with level setting
	fmt.Printf("%s [%s]: %s\n", currTime(), lvl, msg)
}

// currTime returns the current time in the format "Day Mon 02 15:04:05 2006"
func currTime() string {
	return time.Now().Format(time.RFC1123)
}
