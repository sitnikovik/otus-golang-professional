package sender

import (
	"context"
	"fmt"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
)

// Run runs the app.
func (a *App) Run(ctx context.Context) error {
	return a.consumeEvents(ctx)
}

// consumeEvents consumes events from the events queue.
func (a *App) consumeEvents(ctx context.Context) error {
	consumerName := "event_consumer"
	logger.Debugf("consuming events by consumer \"%s\"...", consumerName)

	msgs, err := a.rabbitmq.Consume(ctx, consumerName)
	if err != nil {
		return fmt.Errorf("failed to consume messages: %w", err)
	}

	for msg := range msgs {
		logger.Infof("received event: %s", msg.Body)
	}

	return nil
}
