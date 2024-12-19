package event

import (
	"context"
	"time"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// CreateEvent creates a new event.
func (s *Service) CreateEvent(ctx context.Context, event *eventModel.Event) (uint64, error) {
	logger.Debugf("creating event: %v", *event)

	event.CreatedAt = time.Now()

	id, err := s.db.CreateEvent(ctx, event)
	if err != nil {
		logger.Errorf("failed to create event: %v", err)
		return 0, err
	}

	return id, nil
}
