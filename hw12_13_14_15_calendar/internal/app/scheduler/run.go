package scheduler

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	eventsFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/panics"
)

// Run runs the app.
func (a *App) Run(ctx context.Context) error {
	var wg sync.WaitGroup

	// Run publishing in parallel
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer panics.Recover()

		if err := a.publishEvents(ctx, 1*time.Second); err != nil {
			logger.Criticalf("failed to publish events: %v", err)
		}
	}()

	// Run cleaning in parallel
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer panics.Recover()

		err := a.deleteOldEvents(ctx)
		if err != nil {
			logger.Criticalf("failed to clean old events: %v", err)
		}
	}()

	wg.Wait()
	return nil
}

// publishEvents publishes events to the events queue.
func (a *App) publishEvents(ctx context.Context, interval time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	notNotified := false
	for range ticker.C {
		events, err := a.di.EventService().GetEvents(ctx, eventsFilter.Filter{
			IsNotified: &notNotified,
		})
		if err != nil {
			return fmt.Errorf("failed to get events: %w", err)
		}

		if len(events) == 0 {
			continue
		}

		logger.Infof("found %d events to publish", len(events))
		for _, event := range events {
			eventAsJSON, err := json.Marshal(event)
			if err != nil {
				logger.Errorf("failed to marshal event with id \"%d\": %v", event.ID, err)
				continue
			}

			routingKey := fmt.Sprintf("event.%d", event.ID)
			if err := a.rabbitmq.PublishJSON(ctx, routingKey, eventAsJSON); err != nil {
				logger.Criticalf("failed to publish event with id \"%d\": %v", event.ID, err)
				continue
			}

			logger.Infof("event with id \"%d\" has been published", event.ID)
			event.IsNotified = true
			if err := a.DI().EventService().UpdateEvent(ctx, event); err != nil {
				logger.Alertf("failed to update event with id \"%d\": %v", event.ID, err)
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

	if len(events) == 0 {
		return nil
	}

	// Delete events
	logger.Infof("deleting %d old events...", len(events))
	for _, event := range events {
		if err := a.di.EventService().DeleteEvent(ctx, event.ID); err != nil {
			logger.Criticalf("failed to delete event with id \"%d\": %v", event.ID, err)
			continue
		}
		logger.Debugf("event with id \"%d\" has been deleted", event.ID)
	}

	return nil
}
