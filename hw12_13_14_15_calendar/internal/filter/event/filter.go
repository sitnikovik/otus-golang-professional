package event

import (
	"reflect"
	"time"
)

// Filter describes the filter for the list of events.
type Filter struct {
	// CreatedFrom - the start date of the event creation.
	CreatedFrom *time.Time
	// CreatedTo - the end date of the event creation.
	CreatedTo *time.Time
	// FinishedFrom - the start date of the event finish.
	FinishedFrom *time.Time
	// FinishedTo - the end date of the event finish.
	FinishedTo *time.Time

	// IDs - the list of event IDs.
	IDs []uint64
	// OwnerIDs - the list of event owner IDs.
	OwnerIDs []uint64

	// Limit - the limit of the events.
	Limit uint64
	// Offset - the offset of the events for pagination.
	Offset uint64
}

// Empty checks if the filter is full empty.
func (f Filter) Empty() bool {
	return reflect.DeepEqual(f, Filter{})
}
