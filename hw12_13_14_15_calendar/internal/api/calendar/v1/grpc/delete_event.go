package grpc

import (
	"context"

	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

// DeleteEvent deletes the event.
func (i *Implementation) DeleteEvent(ctx context.Context, req *pkg.DeleteEventRequest) (*pkg.DeleteEventResponse, error) {
	if err := i.eventService.DeleteEvent(ctx, req.GetId()); err != nil {
		return nil, err
	}

	return &pkg.DeleteEventResponse{}, nil
}
