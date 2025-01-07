package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

// GetEventsForMonth returns a list of events that will occur in the current month.
func (i *Implementation) GetEventsForMonth(
	ctx context.Context,
	req *pkg.GetEventsForMonthRequest,
) (*pkg.GetEventsResponse, error) {
	events, err := i.eventService.GetEventsForMonth(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pkg.GetEventsResponse{
		Events: eventsToResponse(events),
		Total:  uint64(len(events)),
	}, nil
}
