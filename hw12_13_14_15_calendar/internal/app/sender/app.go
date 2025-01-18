package sender

import (
	"context"
	"log"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/connections/rabbitmq"
)

// App describes the app instance.
type App struct {
	// config describes the app configuration.
	config config.Config
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
