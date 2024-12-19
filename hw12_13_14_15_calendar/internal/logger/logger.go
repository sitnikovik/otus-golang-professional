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
		log(msg)
	}
}

// Infof logs the info with formatted message
func Infof(format string, a ...interface{}) {
	Info(fmt.Sprintf(format, a...))
}

// Info logs the info message
func Info(msg string) {
	if level >= InfoLevel {
		log(msg)
	}
}

// WarnNoticefingf logs the notice with formatted message
func Noticef(format string, a ...interface{}) {
	Notice(fmt.Sprintf(format, a...))
}

// Notice logs the notice message
func Notice(msg string) {
	if level >= NoticeLevel {
		log(msg)
	}
}

// Warningf logs the warning with formatted message
func Warningf(format string, a ...interface{}) {
	Warning(fmt.Sprintf(format, a...))
}

// Warning logs the warning message
func Warning(msg string) {
	if level >= WarningLevel {
		log(msg)
	}
}

// Errorf logs the error with formatted message
func Errorf(format string, a ...interface{}) {
	Error(fmt.Sprintf(format, a...))
}

// Error logs the error message
func Error(msg string) {
	fmt.Printf("log: %v\n", log)
	if level >= ErrorLevel {
		log(msg)
	}
}

// Alertf logs the alert with formatted message
func Alertf(format string, a ...interface{}) {
	Alert(fmt.Sprintf(format, a...))
}

// Alert logs the alert message
func Alert(msg string) {
	if level >= AlertLevel {
		log(msg)
	}
}

// Critical logs the critical with formatted message
func Criticalf(format string, a ...interface{}) {
	Critical(fmt.Sprintf(format, a...))
}

// Critical logs the critical message
func Critical(msg string) {
	if level >= CriticalLevel {
		log(msg)
	}
}

// Emergency logs the emergency with formatted message
func Emergencyf(format string, a ...interface{}) {
	Emergency(fmt.Sprintf(format, a...))
}

// Emergency logs the emergency message
func Emergency(msg string) {
	if level >= EmergencyLevel {
		log(msg)
		panic(msg)
	}
}

// log logs the message
func log(msg string) {
	// no using slog package cause of diffuculties with level setting
	echo(msg)
}

// echo logs the message with the current time
func echo(msg string) {
	fmt.Printf("%s [%s]: %s\n", currTime(), level, msg)
}

// currTime returns the current time in the format "Day Mon 02 15:04:05 2006"
func currTime() string {
	return time.Now().Format(time.RFC1123)
}
