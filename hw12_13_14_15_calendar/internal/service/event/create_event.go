package event

import (
	"context"
	"fmt"
	"time"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// CreateEvent creates a new event.
func (s *Service) CreateEvent(ctx context.Context, event *eventModel.Event) (uint64, error) {
	logger.Debugf("creating event: %v", *event)

	event.CreatedAt = time.Now()

	if err := validateEventToCreate(event); err != nil {
		logger.Errorf("failed to validate event: %v", err)
		return 0, err
	}

	id, err := s.db.CreateEvent(ctx, event)
	if err != nil {
		logger.Errorf("failed to create event: %v", err)
		return 0, err
	}

	return id, nil
}

// validateEventToCreate validates the event to create.
func validateEventToCreate(event *eventModel.Event) error {
	if event == nil {
		return fmt.Errorf("event is nil")
	}
	if event.Title == "" {
		return fmt.Errorf("event title is empty")
	}
	if event.CreatedAt.IsZero() {
		return fmt.Errorf("event created date at is zero")
	}
	if event.FinishedAt != nil && event.FinishedAt.IsZero() {
		return fmt.Errorf("event finished date at is zero")
	}
	if event.OwnerID == 0 {
		return fmt.Errorf("event owner id is zero")
	}

	return nil
}
