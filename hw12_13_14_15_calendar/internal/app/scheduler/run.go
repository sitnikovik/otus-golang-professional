package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	eventsFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
)

// Run runs the app.
func (a *App) Run(ctx context.Context) error {
	err := a.publishEvents(ctx, 1*time.Second)
	if err != nil {
		return fmt.Errorf("failed to publish events: %w", err)
	}

	err = a.deleteOldEvents(ctx)
	if err != nil {
		return fmt.Errorf("failed to clean old events: %w", err)
	}

	return nil
}

// publishEvents publishes events to the events queue.
func (a *App) publishEvents(ctx context.Context, interval time.Duration) error {
	ticker := time.NewTicker(interval) // TODO: move to config
	defer ticker.Stop()

	for range ticker.C {
		events, err := a.di.EventService().GetEventsForToday(ctx)
		if err != nil {
			return fmt.Errorf("failed to get events: %w", err)
		}

		for _, event := range events {
			eventAsJSON, err := json.Marshal(event)
			if err != nil {
				return fmt.Errorf("failed to marshal event with id \"%d\": %w", event.ID, err)
			}

			// TODO: add routing key
			if err := a.rabbitmq.PublishJSON(ctx, "", eventAsJSON); err != nil {
				return fmt.Errorf("failed to publish message: %w", err)
			}
		}
	}

	return nil
}

// deleteOldEvents deletes events old events to free up the storage.
func (a *App) deleteOldEvents(ctx context.Context) error {
	// Get events for the past year
	createdTo := time.Now().AddDate(-1, 0, 0)
	events, err := a.di.EventService().GetEvents(ctx, eventsFilter.Filter{
		CreatedTo: &createdTo,
	})
	if err != nil {
		return fmt.Errorf("failed to get events: %w", err)
	}

	// Delete events
	for _, event := range events {
		if err := a.di.EventService().DeleteEvent(ctx, event.ID); err != nil {
			return fmt.Errorf("failed to delete event with id \"%d\": %w", event.ID, err)
		}
	}

	return nil
}
