package grpc

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/api/calendar/v1/grpc/interceptor"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	eventFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

// eventService describes the event service.
type eventService interface {
	// CreateEvent creates a new event.
	CreateEvent(ctx context.Context, event *eventModel.Event) (uint64, error)
	// UpdateEvent updates the event.
	UpdateEvent(ctx context.Context, event *eventModel.Event) error
	// DeleteEvent deletes the event.
	DeleteEvent(ctx context.Context, eventID uint64) error
	// GetEvent returns the event by ID.
	GetEvent(ctx context.Context, eventID uint64) (*eventModel.Event, error)
	// GetEvents returns the events by filter.
	GetEvents(ctx context.Context, filter eventFilter.Filter) ([]*eventModel.Event, error)
}

// implementation describes the gRPC server implementation.
type Implementation struct {
	pkg.UnimplementedEventServiceServer

	// eventService describes the event service instance.
	eventService eventService
}

// Server describes the app gRPC server implementation.
type Server struct {
	// server - gRPC server instance.
	server *grpc.Server
	// config - the gRPC server configuration.
	config config.GRPCConf
}

// NewServer creates a new gRPC server.
func NewServer(conf config.GRPCConf, eventService eventService) *Server {
	s := &Server{
		config: conf,
		server: grpc.NewServer(
			grpc.UnaryInterceptor(interceptor.Logging),
		),
	}

	pkg.RegisterEventServiceServer(s.server, newImplementation(eventService))

	return s
}

// newImplementation creates and returns a new gRPC server implementation.
func newImplementation(eventService eventService) *Implementation {
	return &Implementation{
		eventService: eventService,
	}
}

// Serve starts the gRPC server.
func (s *Server) Serve(ctx context.Context) error {
	address := s.config.Host + ":" + s.config.Port
	logger.Infof("Starting gRPC server on %s", address)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("failed to listen gRPC server: %w", err)
	}

	if err = s.server.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve gRPC server: %w", err)
	}

	// Wait for the context to be done
	<-ctx.Done()

	s.GracefulShutdown(ctx)

	return nil
}

// GracefulShutdown stops the gRPC server.
func (s *Server) GracefulShutdown(_ context.Context) {
	logger.Infof("Stopping gRPC server")
	s.server.GracefulStop()
}
