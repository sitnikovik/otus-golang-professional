package main

import (
	"context"
	"flag"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	calendarHttpServer "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/server/http/calendar"
)

var configPath string

func init() {
	flag.StringVar(
		&configPath,
		"config",
		".env",
		"Path to configuration file",
	)
}

func main() {
	// Cmd parsing
	flag.Parse()
	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	// Configuration init
	config, err := config.NewConfig(configPath)
	if err != nil {
		logger.Emergencyf("failed to load config: %v", err)
	}
	logger.Debugf("Used config file: %s", configPath)

	configLevel := config.Logger.Level
	logger.SetLevel(logger.LevelFromString(configLevel))
	logger.Debugf("Specified log level: %s", configLevel)

	// App init
	// di := depinjection.NewDIContainer(config)
	calendarApp := app.New(
		app.NewDIContainer(config),
	)

	// Servers
	httpServer := calendarHttpServer.NewServer(calendarApp, config.HTTP)
	grpcServer := grpc.NewServer()

	// Run the app
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP,
	)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// Graceful shutdown
		defer wg.Done()
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		grpcServer.GracefulStop()

		if err := httpServer.Stop(ctx); err != nil {
			logger.Criticalf("failed to stop http server: %v", err)
		}
	}()
	// Start the server
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Start http server
		if err := httpServer.Start(ctx, config.HTTP); err != nil {
			logger.Errorf("failed to start http server: %v", err)
			cancel()
			os.Exit(1)
		}

	}()
	// Start the gRPC server
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Start gRPC server
		address := config.GRPC.Host + ":" + config.GRPC.Port
		logger.Infof("Starting gRPC server on %s", address)
		grpcLis, err := net.Listen("tcp", address)
		if err != nil {
			logger.Errorf("failed to listen gRPC server: %v", err)
			cancel()
			os.Exit(1)
		}
		if err = grpcServer.Serve(grpcLis); err != nil {
			logger.Errorf("failed to serve gRPC server: %v", err)
			cancel()
			os.Exit(1)
		}
	}()
	wg.Wait()
}
