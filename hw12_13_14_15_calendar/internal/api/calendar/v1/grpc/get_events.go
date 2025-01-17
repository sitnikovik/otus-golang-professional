package grpc

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	eventFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

func (i *Implementation) GetEvents(ctx context.Context, req *pkg.GetEventsRequest) (*pkg.GetEventsResponse, error) {
	filter := eventsRequestToFilter(req)
	events, err := i.eventService.GetEvents(ctx, filter)
	if err != nil {
		return nil, err
	}

	return &pkg.GetEventsResponse{
		Events: eventsToResponse(events),
		Total:  uint64(len(events)),
	}, nil
}

// eventsRequestToFilter converts the GetEventsRequest to the event filter.
func eventsRequestToFilter(req *pkg.GetEventsRequest) eventFilter.Filter {
	if req == nil {
		return eventFilter.Filter{}
	}

	filter := eventFilter.Filter{}
	if req.CreatedFrom != nil {
		t := FromGRPCTime(req.CreatedFrom)
		filter.CreatedFrom = &t
	}
	if req.CreatedTo != nil {
		t := FromGRPCTime(req.CreatedTo)
		filter.CreatedTo = &t
	}
	if req.FinishedFrom != nil {
		t := FromGRPCTime(req.FinishedFrom)
		filter.FinishedFrom = &t
	}
	if req.FinishedTo != nil {
		t := FromGRPCTime(req.FinishedTo)
		filter.FinishedTo = &t
	}
	if len(req.Ids) != 0 {
		filter.IDs = req.Ids
	}
	if req.OwnerId != 0 {
		filter.OwnerIDs = append(filter.OwnerIDs, req.OwnerId)
	}
	if req.Limit > 0 {
		filter.Limit = req.Limit
	}
	if req.Offset > 0 {
		filter.Offset = req.Offset
	}

	return filter
}

// eventsToResponse converts the list of events to the list of events in the response.
func eventsToResponse(events []*event.Event) []*pkg.Event {
	var eventsList []*pkg.Event
	for _, event := range events {
		var finishedAt *timestamppb.Timestamp
		if event.FinishedAt != nil {
			finishedAt = ToGRPCTime(*event.FinishedAt)
		}
		eventsList = append(eventsList, &pkg.Event{
			Id:           event.ID,
			Title:        event.Title,
			Description:  event.Description,
			CreatedAt:    ToGRPCTime(event.CreatedAt),
			FinishedAt:   finishedAt,
			OwnerId:      event.OwnerID,
			NotifyBefore: int64(event.NotifyBefore),
			IsNotified:   event.IsNotified,
		})
	}

	return eventsList
}
