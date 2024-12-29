package event

import (
	"context"
	"time"

	eventFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
	timeUtils "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/utils/time"
)

// GetEventsBeforeDays returns a list of events that will occur in the next N days.
func (s *Service) GetEventsBeforeDays(ctx context.Context, days uint32) ([]*eventModel.Event, error) {
	logger.Debugf("getting events that will occur in the next %d days", days)

	timeNow := time.Now()
	finishedTo := timeUtils.UpToDays(timeNow, int(days))

	events, err := s.GetEvents(ctx, eventFilter.Filter{
		CreatedFrom:  &timeNow,
		FinishedFrom: &timeNow,
		FinishedTo:   &finishedTo,
	})
	if err != nil {
		logger.Errorf("failed to get events: %v", err)
		return nil, err
	}

	return events, nil
}
