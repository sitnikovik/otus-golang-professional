package event

import (
	"context"
	"time"

	eventFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
	timeUtils "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/utils/time"
)

// GetEventsForWeek returns a list of events that will occur in the current week.
func (s *Service) GetEventsForWeek(ctx context.Context) ([]*eventModel.Event, error) {
	logger.Debug("getting events for current week")

	fromTime := time.Now()
	toTime := timeUtils.EndOfWeek(fromTime)

	events, err := s.GetEvents(ctx, eventFilter.Filter{
		CreatedFrom:  &fromTime,
		FinishedFrom: &fromTime,
		FinishedTo:   &toTime,
	})
	if err != nil {
		logger.Errorf("failed to get events: %v", err)
		return nil, err
	}

	return events, nil
}
