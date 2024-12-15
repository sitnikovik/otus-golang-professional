package app

import (
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app/depinjection"
)

type App struct { // TODO
	di depinjection.DIContainer
}

// DIContainer returns the DI container instance
func (a *App) DIContainer() depinjection.DIContainer {
	return nil
}

// New creates and returns the app instance
func New(di depinjection.DIContainer) *App {
	return &App{
		di: di,
	}
}
