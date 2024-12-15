package pgsql

import (
	"context"

	"github.com/Masterminds/squirrel"

	eventFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// allEventColumns is the list of fields for the event entity
var allEventColumns = []string{
	"id",
	"title",
	"created_at",
	"finished_at",
	"description",
	"owner_id",
	"notify_before",
}

// GetEvents returns the events by filter
func (s *PgStorage) GetEvents(ctx context.Context, filter eventFilter.Filter) ([]*eventModel.Event, error) {
	sb := squirrel.
		Select(allEventColumns...).
		From(eventsTable).
		PlaceholderFormat(squirrel.Dollar)

	if len(filter.IDs) > 0 {
		sb = sb.Where(squirrel.Eq{"id": filter.IDs})
	}
	if filter.Limit > 0 {
		sb = sb.Limit(filter.Limit)
	}

	sql, args, err := sb.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.Query(sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*eventModel.Event
	for rows.Next() {
		var event eventModel.Event
		if err := rows.Scan(&event.ID, &event.Title); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	return events, nil
}
