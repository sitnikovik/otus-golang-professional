package sender

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
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
		var event eventModel.Event
		if err := json.Unmarshal(msg.Body, &event); err != nil {
			logger.Errorf("failed to unmarshal event: %v", err)
			continue
		}

		event.IsNotified = true
		if err := a.DI().EventService().UpdateEvent(ctx, &event); err != nil {
			logger.Criticalf("failed to update event with id %d: %v", event.ID, err)
			continue
		}
	}

	return nil
}
