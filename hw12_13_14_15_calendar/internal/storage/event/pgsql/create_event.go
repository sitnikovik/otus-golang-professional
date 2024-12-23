package pgsql

import (
	"context"

	"github.com/Masterminds/squirrel"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// CreateEvent creates a new event.
func (s *PgStorage) CreateEvent(_ context.Context, event *eventModel.Event) (uint64, error) {
	sb := squirrel.
		Insert(eventsTable).
		PlaceholderFormat(squirrel.Dollar).
		Columns(eventColumnsToInsert...).
		Values(s.parseEventValuesForColumns(event)...).
		Suffix("RETURNING id")

	sql, args, err := sb.ToSql()
	if err != nil {
		return 0, err
	}

	rows, err := s.db.Query(sql, args...)
	if err != nil {
		return 0, err
	}

	var id uint64
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
	}

	return id, nil
}

func (s *PgStorage) parseEventValuesForColumns(event *eventModel.Event) []interface{} {
	if event == nil {
		return nil
	}

	values := make([]interface{}, 0, len(eventColumnsToInsert))
	for _, c := range eventColumnsToInsert {
		switch c {
		case "title":
			values = append(values, event.Title)
		case "created_at":
			values = append(values, event.CreatedAt)
		case "finished_at":
			values = append(values, event.FinishedAt)
		case "description":
			values = append(values, event.Description)
		case "owner_id":
			values = append(values, event.OwnerID)
		case "notify_before":
			values = append(values, event.NotifyBefore)
		}
	}

	return values
}
