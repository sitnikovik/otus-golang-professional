package grpc

import (
	"context"

	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

// GetEvent returns the event by ID.
func (i *Implementation) GetEvent(ctx context.Context, req *pkg.GetEventRequest) (*pkg.GetEventResponse, error) {
	event, err := i.eventService.GetEvent(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pkg.GetEventResponse{
		Event: &pkg.Event{
			Id:           event.ID,
			Title:        event.Title,
			Description:  event.Description,
			CreatedAt:    ToGRPCTime(event.CreatedAt),
			FinishedAt:   ToGRPCTime(*event.FinishedAt),
			OwnerId:      event.OwnerID,
			NotifyBefore: int64(event.NotifyBefore),
		},
	}, nil
}
