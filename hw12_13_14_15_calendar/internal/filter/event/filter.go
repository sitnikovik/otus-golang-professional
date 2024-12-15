package event

import "reflect"

// Filter describes the filter for the list of events
type Filter struct {
	IDs   []string
	Limit uint64
}

// Empty checks if the filter is full empty
func (f Filter) Empty() bool {
	return reflect.DeepEqual(f, Filter{})
}
