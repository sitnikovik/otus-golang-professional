package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

// GetEventsBeforeDays returns the events before days.
func (i *Implementation) GetEventsBeforeDays(
	ctx context.Context,
	req *pkg.GetEventsBeforeDaysRequest,
) (*pkg.GetEventsResponse, error) {
	events, err := i.eventService.GetEventsBeforeDays(ctx, req.GetDays())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pkg.GetEventsResponse{
		Events: eventsToResponse(events),
		Total:  uint64(len(events)),
	}, nil
}
