package app

import (
	"context"
	"strconv"

	"github.com/jackc/pgx"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/connections/pg"
	eventFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
	events "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/service/event"
	eventStorageSql "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/storage/event/pgsql"
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
	// GetEventsBeforeDays returns the events before days.
	GetEventsBeforeDays(ctx context.Context, days uint32) ([]*eventModel.Event, error)
	// GetEventsForMonth returns a list of events that will occur in the current month.
	GetEventsForMonth(ctx context.Context) ([]*eventModel.Event, error)
	// GetEventsForWeek returns a list of events that will occur in the current week.
	GetEventsForWeek(ctx context.Context) ([]*eventModel.Event, error)
	// GetEventsForToday returns a list of events that will occur today.
	GetEventsForToday(ctx context.Context) ([]*eventModel.Event, error)
}

// DIContainer describes the DI container instance.
type DIContainer struct {
	conf config.Config

	eventService eventService

	pgx *pgx.ConnPool
}

// NewDIContainer creates and returns the DI container instance.
func NewDIContainer(conf config.Config) *DIContainer {
	di := &DIContainer{
		conf: conf,
	}

	return di
}

// EventService returns the event service instance.
func (d *DIContainer) EventService() eventService {
	if d.eventService == nil {
		d.eventService = events.NewService(
			eventStorageSql.New(
				d.pg(),
			),
		)
	}

	return d.eventService
}

// Postgres returns the postgres storage instance.
func (d *DIContainer) pg() *pgx.ConnPool {
	if d.pgx == nil {
		pgPort, _ := strconv.Atoi(d.conf.PG.Port)
		pgx, err := pg.NewConnPool(
			d.conf.PG.Database,
			d.conf.PG.User,
			d.conf.PG.Password,
			d.conf.PG.Host,
			pgPort,
		)
		if err != nil {
			logger.Panicf("failed to connect to postgres: %v", err)
		}
		d.pgx = pgx
	}

	return d.pgx
}
