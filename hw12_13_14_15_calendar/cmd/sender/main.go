package main

import (
	"context"
	"flag"

	senderApp "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app/sender"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
)

// configPath is the path to the configuration file to specified by the user via flag.
var configPath string

func init() {
	flag.StringVar(
		&configPath,
		"config",
		"/etc/sender/.env",
		"Path to configuration file",
	)
}

func main() {
	// Configuration init
	config := initConfig()

	// App init
	ctx := context.Background()
	senderApp := senderApp.NewApp(ctx, config)

	// Run the app
	if err := senderApp.Run(ctx); err != nil {
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

	configLevel := config.Logger.Level
	logger.SetLevel(logger.LevelFromString(configLevel))
	logger.Debugf("Specified log level: %s", configLevel)

	return config
}
