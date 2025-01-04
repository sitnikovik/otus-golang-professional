package grpc

import (
	"context"
	"time"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

// UpdateEvent updates the event.
func (i *Implementation) UpdateEvent(ctx context.Context, req *pkg.UpdateEventRequest) (*pkg.UpdateEventResponse, error) {
	finishedAt := FromGRPCTime(req.Event.FinishedAt)
	e := &event.Event{
		ID:           req.Event.Id,
		Title:        req.Event.Title,
		Description:  req.Event.Description,
		FinishedAt:   &finishedAt,
		OwnerID:      req.Event.OwnerId,
		NotifyBefore: time.Duration(req.Event.NotifyBefore),
	}

	if err := i.eventService.UpdateEvent(ctx, e); err != nil {
		return nil, err
	}

	return &pkg.UpdateEventResponse{}, nil
}
