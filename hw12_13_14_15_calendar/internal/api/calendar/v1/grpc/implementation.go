package grpc

import (
	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

// implementation describes the gRPC server implementation.
type implementation struct {
	pkg.UnimplementedEventServiceServer

	// eventService describes the event service instance.
	eventService eventService
}

// newImplementation creates and returns a new gRPC server implementation.
func newImplementation(eventService eventService) *implementation {
	return &implementation{
		eventService: eventService,
	}
}
