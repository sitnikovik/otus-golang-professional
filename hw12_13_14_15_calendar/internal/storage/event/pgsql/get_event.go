package pgsql

import (
	"context"
	"fmt"

	eventFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// GetEvent returns the event by ID
func (s *PgStorage) GetEvent(ctx context.Context, eventID string) (*eventModel.Event, error) {
	filter := eventFilter.Filter{
		IDs:   []string{eventID},
		Limit: 1,
	}

	events, err := s.GetEvents(ctx, filter)
	if err != nil {
		return nil, err
	}
	if len(events) == 0 {
		return nil, fmt.Errorf("event with id(%s) not found", eventID)
	}

	return events[0], nil
}
