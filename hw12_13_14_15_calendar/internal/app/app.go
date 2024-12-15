package app

type App struct {
	di *DIContainer
}

// New creates and returns the app instance
func New(di *DIContainer) *App {
	return &App{
		di: di,
	}
}

func (a *App) DI() *DIContainer {
	return a.di
}
