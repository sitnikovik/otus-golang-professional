package panics

import "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"

// Recover recovers from panic.
func Recover() {
	if r := recover(); r != nil {
		logger.Emergencyf("recovered from panic: %v", r)
	}
}
