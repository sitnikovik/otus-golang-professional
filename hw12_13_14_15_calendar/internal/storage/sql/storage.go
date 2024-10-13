package sqlstorage

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/logger"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/storage"
)

const (
	// eventsTable is the name of the table for events
	eventsTable = "events"
)

var (
	// allEventColumns is the list of fields for the event entity
	allEventColumns = []string{
		"id",
		"title",
		"created_at",
		"finished_at",
		"description",
		"owner_id",
		"notify_before",
	}
)

// ListFilter describes the filter for the list of events
type ListFilter struct {
	IDs []string
}

type Storage interface {
	// CreateEvent creates a new event
	CreateEvent(ctx context.Context, event *storage.Event) (string, error)
	// UpdateEvent updates the event
	UpdateEvent(ctx context.Context, event *storage.Event) error
	// DeleteEvent deletes the event
	DeleteEvent(ctx context.Context, eventID string) error
	// GetEvent returns the event by ID
	GetEvent(ctx context.Context, eventID string) (*storage.Event, error)
	// GetEvents returns the events by filter
	GetEvents(ctx context.Context, filter ListFilter) ([]*storage.Event, error)

	// Close closes the storage
	Close(ctx context.Context) error
}

type pgStorage struct {
	db *pgx.Conn // Пул коннектов к БД
}

// New creates and returns the sql storage instance
func New(pg *pgx.Conn) Storage {
	return &pgStorage{
		db: pg,
	}
}

// Close closes the storage
func (s *pgStorage) Close(ctx context.Context) error {
	return s.db.Close(ctx)
}

// CreateEvent creates a new event
func (s *pgStorage) CreateEvent(ctx context.Context, event *storage.Event) (string, error) {
	logger.Debugf("creating event: %v", *event)

	sb := squirrel.
		Insert(eventsTable).
		PlaceholderFormat(squirrel.Dollar).
		Columns("id", "title").
		Values(event.ID, event.Title).
		Suffix("RETURNING id")

	sql, args, err := sb.ToSql()
	if err != nil {
		return "", err
	}

	rows, err := s.db.Query(ctx, sql, args...)
	if err != nil {
		return "", err
	}

	var id string
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return "", err
		}
	}

	return id, nil
}

// UpdateEvent updates the event
func (s *pgStorage) UpdateEvent(ctx context.Context, event *storage.Event) error {
	logger.Debugf("updating event with id(%s)", event.ID)

	sb := squirrel.
		Update(eventsTable).
		PlaceholderFormat(squirrel.Dollar).
		Set("title", event.Title).
		Where(squirrel.Eq{"id": event.ID})

	sql, args, err := sb.ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

// DeleteEvent deletes the event
func (s *pgStorage) DeleteEvent(ctx context.Context, eventID string) error {
	logger.Infof("deleting event by id(%s)", eventID)

	sb := squirrel.
		Delete(eventsTable).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": eventID})

	sql, args, err := sb.ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}

	return nil
}

// GetEvent returns the event by ID
func (s *pgStorage) GetEvent(ctx context.Context, eventID string) (*storage.Event, error) {
	logger.Debugf("getting event by id(%s)", eventID)

	sb := squirrel.
		Select(allEventColumns...).
		From(eventsTable).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": eventID})

	sql, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	row := s.db.QueryRow(ctx, sql, args...)

	var event storage.Event
	if err := row.Scan(&event.ID, &event.Title); err != nil {
		return nil, err
	}

	return &event, nil
}

// GetEvents returns the events by filter
func (s *pgStorage) GetEvents(ctx context.Context, filter ListFilter) ([]*storage.Event, error) {
	sb := squirrel.
		Select(allEventColumns...).
		From(eventsTable).
		PlaceholderFormat(squirrel.Dollar)

	if len(filter.IDs) > 0 {
		sb = sb.Where(squirrel.Eq{"id": filter.IDs})
	}

	sql, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*storage.Event
	for rows.Next() {
		var event storage.Event
		if err := rows.Scan(&event.ID, &event.Title); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	return events, nil
}
