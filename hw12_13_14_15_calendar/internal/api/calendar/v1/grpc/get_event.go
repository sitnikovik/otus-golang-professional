package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

// GetEvent returns the event by ID.
func (i *Implementation) GetEvent(ctx context.Context, req *pkg.GetEventRequest) (*pkg.GetEventResponse, error) {
	event, err := i.eventService.GetEvent(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	var finishedAt *timestamppb.Timestamp
	if event.FinishedAt != nil {
		finishedAt = ToGRPCTime(*event.FinishedAt)
	}

	return &pkg.GetEventResponse{
		Event: &pkg.Event{
			Id:           event.ID,
			Title:        event.Title,
			Description:  event.Description,
			CreatedAt:    ToGRPCTime(event.CreatedAt),
			FinishedAt:   finishedAt,
			OwnerId:      event.OwnerID,
			NotifyBefore: int64(event.NotifyBefore),
		},
	}, nil
}
