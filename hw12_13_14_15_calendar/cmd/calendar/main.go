package main

import (
	"context"
	"flag"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
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
	config := initConfig()

	// App init
	ctx := context.Background()
	calendarApp := app.New(ctx, config)

	// Run the app
	if err := calendarApp.Run(); err != nil {
		logger.Panicf("failed to run app: %v", err)
	}
}

// initConfig initializes the app configuration.
func initConfig() config.Config {
	// Configuration init
	logger.Infof("Loading config from file: %s", configPath)
	config, err := config.NewConfig(configPath)
	if err != nil {
		logger.Panicf("failed to load config: %v", err)
	}
	logger.Debugf("Used config file: %s", configPath)

	configLevel := config.Logger.Level
	logger.SetLevel(logger.LevelFromString(configLevel))
	logger.Debugf("Specified log level: %s", configLevel)

	return config
}
