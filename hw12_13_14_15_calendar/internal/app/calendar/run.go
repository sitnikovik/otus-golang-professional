package calendar

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/api/calendar/v1/grpc"
	calendarHttpServer "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/api/calendar/v1/http"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app/panics"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
)

// Run runs the app.
func (a *App) Run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(
		ctx,
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP,
	)
	defer cancel()

	httpAppServer := calendarHttpServer.NewServer(a.config.HTTP, a.DI().EventService())
	grpcAppServer := grpc.NewServer(a.config.GRPC, a.DI().EventService())

	var wg sync.WaitGroup

	// Graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-sigChan

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		logger.Infof("Graceful shutdown started...")
		a.GracefulShutdown(ctx)
	}()

	// Start the server
	wg.Add(1)
	go func() {
		defer panics.Recover()
		defer wg.Done()
		// Start http server
		if err := a.runHTTPServer(ctx, httpAppServer); err != nil {
			cancel()
		}
	}()

	// Start the gRPC server
	wg.Add(1)
	go func() {
		defer panics.Recover()
		defer wg.Done()
		if err := a.runGRPCServer(ctx, grpcAppServer); err != nil {
			cancel()
		}
	}()

	wg.Wait()
	logger.Infof("Application stopped")
	return nil
}

// runHTTPServer runs the HTTP server.
func (a *App) runHTTPServer(ctx context.Context, server *calendarHttpServer.Server) error {
	err := server.Serve(ctx)
	if err != nil {
		logger.Emergencyf("failed to run HTTP server: %v", err)
	}

	return err
}

// runGRPCServer runs the gRPC server.
func (a *App) runGRPCServer(ctx context.Context, server *grpc.Server) error {
	err := server.Serve(ctx)
	if err != nil {
		logger.Emergencyf("failed to run gRPC server: %v", err)
	}

	return err
}
