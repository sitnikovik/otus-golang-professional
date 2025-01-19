package sender

import (
	"context"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/app"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/connections/rabbitmq"
)

// init initializes the app dependencies.
func (a *App) init(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initDI,
		a.initRabbitMQ,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

// initDI initializes the DI container.
func (a *App) initDI(_ context.Context) error {
	a.di = app.NewDIContainer(a.config)

	return nil
}

// initRabbitMQ initializes the RabbitMQ connection.
func (a *App) initRabbitMQ(_ context.Context) error {
	dsn := rabbitmq.NewDSN(
		a.config.RabbitMQ.Host,
		a.config.RabbitMQ.Port,
		a.config.RabbitMQ.User,
		a.config.RabbitMQ.Password,
	)
	conn, err := rabbitmq.NewRabbitMQ(
		dsn,
		a.config.EventsQueueName,
		a.config.EventsExchangeName,
		a.config.EventsExchangeType,
	)
	if err != nil {
		return err
	}

	a.rabbitmq = conn

	return nil
}
