package calendar

import (
	"context"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/api/calendar/v1/grpc"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/api/calendar/v1/http"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app"
)

func (a *App) init(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDI,
		a.initGRPCServer,
		a.initHTTPServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

// initDI initializes the DI container.
func (a *App) initDI(ctx context.Context) error {
	a.di = app.NewDIContainer(a.config)

	return nil
}

// initGRPCServer initializes the gRPC server.
func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(a.config.GRPC, a.di.EventService())

	return nil
}

// initHTTPServer initializes the HTTP server.
func (a *App) initHTTPServer(ctx context.Context) error {
	a.httpServer = http.NewServer(a.config.HTTP, a.di.EventService())

	return nil
}
