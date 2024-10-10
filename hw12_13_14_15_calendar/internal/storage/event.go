package storage

import "errors"

var (
	// ErrNotFound is used when the requested event is not found
	ErrNotFound = errors.New("event not found")
)

type Event struct {
	ID    string
	Title string
	// TODO
}
