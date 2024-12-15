package event

import (
	"context"

	eventFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// eventDB is the interface to work with events db storage.
type eventDB interface {
	// CreateEvent creates a new event.
	CreateEvent(ctx context.Context, event *eventModel.Event) (string, error)
	// UpdateEvent updates the event.
	UpdateEvent(ctx context.Context, event *eventModel.Event) error
	// DeleteEvent deletes the event.
	DeleteEvent(ctx context.Context, eventID string) error
	// GetEvent returns the event by ID.
	GetEvent(ctx context.Context, eventID string) (*eventModel.Event, error)
	// GetEvents returns the events by filter.
	GetEvents(ctx context.Context, filter eventFilter.Filter) ([]*eventModel.Event, error)

	// Close closes the storage
	Close(ctx context.Context) error
}

// Service is the service to work with events.
type Service struct {
	// db is the events db storage.
	db eventDB
}

// NewService creates and returns the service instance.
func NewService(db eventDB) *Service {
	return &Service{
		db: db,
	}
}
