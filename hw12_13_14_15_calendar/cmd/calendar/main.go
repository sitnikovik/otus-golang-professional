package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"sync"
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

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logger.Criticalf("failed to stop http server: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.Start(ctx); err != nil {
			logger.Errorf("failed to start http server: %v", err)
			cancel()
			os.Exit(1) //nolint:gocritic
		}
	}()

	wg.Wait()

}
