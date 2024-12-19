package event

import (
	"context"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// UpdateEvent updates the event.
func (s *Service) UpdateEvent(ctx context.Context, event *eventModel.Event) error {
	logger.Debugf("updating event with id(%d)", event.ID)

	if err := s.db.UpdateEvent(ctx, event); err != nil {
		logger.Errorf("failed to update event: %v", err)
		return err
	}

	return nil
}
