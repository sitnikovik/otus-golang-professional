package event

import (
	"context"
	"fmt"

	eventFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// GetEvents returns the events by filter.
func (s *Service) GetEvents(ctx context.Context, filter eventFilter.Filter) ([]*eventModel.Event, error) {
	logMsg := "getting events by filter: "
	if filter.Empty() {
		logMsg = "getting all events cause filter is empty"
	} else {
		if len(filter.IDs) > 0 {
			logMsg += fmt.Sprintf("ids(%v)", filter.IDs)
		}
		if filter.Limit > 0 {
			logMsg += fmt.Sprintf(" limit(%d)", filter.Limit)
		}
	}
	logger.Debug(logMsg)

	return s.db.GetEvents(ctx, filter)
}
