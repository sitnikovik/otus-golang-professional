package main

import (
	"context"

	schedulerApp "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app/scheduler"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
)

// configPath is the path to the configuration file to specified by the user via flag.
var configPath string

func main() {
	// Configuration init
	config := initConfig()

	// App init
	ctx := context.Background()
	schedulerApp := schedulerApp.NewApp(ctx, config)

	// Run the app
	if err := schedulerApp.Run(ctx); err != nil {
		logger.Panicf("failed to run app: %v", err)
	}
}

// initConfig initializes the app configuration.
func initConfig() config.Config {
	// Configuration init
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
