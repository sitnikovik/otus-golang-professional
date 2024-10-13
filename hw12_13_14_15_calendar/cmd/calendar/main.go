package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app/depinjection"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/server/http"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", ".env", "Path to configuration file")
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
		log.Fatal(fmt.Errorf("failed to load config: %w", err))
	}
	logger.SetLevel(logger.LevelFromString(config.Logger.Level))

	// App init
	di := depinjection.NewDIContainer(config)
	calendarApp := app.New(di)

	// Servers
	server := internalhttp.NewServer(calendarApp, config.HTTP)

	// Run the app
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP,
	)
	defer cancel()
	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logger.Error("failed to stop http server: " + err.Error())
		}
	}()
	logger.Info("calendar is running...")
	if err := server.Start(ctx); err != nil {
		logger.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
