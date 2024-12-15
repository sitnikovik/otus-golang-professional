package pgsql

import (
	"context"

	"github.com/Masterminds/squirrel"

	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// CreateEvent creates a new event
func (s *PgStorage) CreateEvent(_ context.Context, event *eventModel.Event) (string, error) {
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

	rows, err := s.db.Query(sql, args...)
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
