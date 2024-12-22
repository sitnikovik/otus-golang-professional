package event

import (
	"context"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// GetEvent returns the event by ID.
func (s *Service) GetEvent(ctx context.Context, eventID uint64) (*eventModel.Event, error) {
	logger.Debugf("getting event by id(%d)", eventID)

	event, err := s.db.GetEvent(ctx, eventID)
	if err != nil {
		logger.Errorf("failed to get event by id(%d): %v", eventID, err)
		return nil, err
	}

	return event, nil
}
