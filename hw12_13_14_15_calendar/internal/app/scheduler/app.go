package scheduler

import (
	"context"
	"log"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/connections/rabbitmq"
)

type App struct {
	// config describes the app configuration.
	config config.Config
	// di describes the DI container instance to store the app dependencies.
	di *app.DIContainer
	// rabbitMQ describes the RabbitMQ connection instance.
	rabbitmq *rabbitmq.RabbitMQ
}

// New creates and returns the app instance.
func NewApp(ctx context.Context, config config.Config) *App {
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
