package grpc

import (
	"context"
	"time"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

// CreateEvent creates a new event.
func (i *Implementation) CreateEvent(ctx context.Context, req *pkg.CreateEventRequest) (*pkg.CreateEventResponse, error) {
	finishedAt := FromGRPCTime(req.FinishedAt)
	e := &event.Event{
		Title:        req.Title,
		Description:  req.Description,
		FinishedAt:   &finishedAt,
		OwnerID:      req.OwnerId,
		NotifyBefore: time.Duration(req.NotifyBefore),
	}

	id, err := i.eventService.CreateEvent(ctx, e)
	if err != nil {
		return nil, err
	}

	return &pkg.CreateEventResponse{
		Id: id,
	}, nil
}
