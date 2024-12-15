package event

import (
	"context"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
)

// DeleteEvent deletes the event.
func (s *Service) DeleteEvent(ctx context.Context, eventID string) error {
	logger.Debugf("deleting event by id(%s)", eventID)

	if err := s.db.DeleteEvent(ctx, eventID); err != nil {
		logger.Errorf("failed to delete event: %v", err)
		return err
	}

	return nil
}
