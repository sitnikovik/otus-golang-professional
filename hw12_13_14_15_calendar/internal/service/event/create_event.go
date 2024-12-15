package event

import (
	"context"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// CreateEvent creates a new event.
func (s *Service) CreateEvent(ctx context.Context, event *eventModel.Event) (string, error) {
	logger.Debugf("creating event: %v", *event)

	id, err := s.db.CreateEvent(ctx, event)
	if err != nil {
		logger.Errorf("failed to create event: %v", err)
		return "", err
	}

	return id, nil
}
