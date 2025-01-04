package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

// GetEventsForToday returns a list of events that will occur today.
func (i *Implementation) GetEventsForToday(
	ctx context.Context,
	req *pkg.GetEventsForTodayRequest,
) (*pkg.GetEventsResponse, error) {
	events, err := i.eventService.GetEventsForToday(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pkg.GetEventsResponse{
		Events: eventsToResponse(events),
		Total:  uint64(len(events)),
	}, nil
}
