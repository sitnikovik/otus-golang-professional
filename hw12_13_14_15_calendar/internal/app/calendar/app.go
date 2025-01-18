package calendar

import (
	"context"
	"log"
	"sync"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/api/calendar/v1/grpc"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/api/calendar/v1/http"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
)

type App struct {
	// config describes the app configuration.
	config config.Config
	// di describes the DI container instance to store the app dependencies.
	di *app.DIContainer
	// grpcServer describes the gRPC server instance.
	grpcServer *grpc.Server
	// httpServer describes the HTTP server instance.
	httpServer *http.Server
}

// New creates and returns the app instance.
func New(ctx context.Context, config config.Config) *App {
	a := &App{
		config: config,
	}

	if err := a.init(ctx); err != nil {
		log.Panicf("failed to initialize the app: %v", err)
	}

	return a
}

// DI returns the DI container instance to use the app dependencies.
func (a *App) DI() *app.DIContainer {
	return a.di
}

// GracefulShutdown performs the graceful shutdown of the app.
func (a *App) GracefulShutdown(ctx context.Context) error {
	shutdowns := []func(context.Context){
		a.grpcServer.GracefulShutdown,
		a.httpServer.GracefulShutdown,
	}

	var wg sync.WaitGroup
	for _, f := range shutdowns {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f(ctx)
		}()
	}
	wg.Wait()

	return nil
}
