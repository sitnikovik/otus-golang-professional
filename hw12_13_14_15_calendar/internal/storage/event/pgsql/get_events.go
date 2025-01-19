package pgsql

import (
	"context"

	"github.com/Masterminds/squirrel"

	eventFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

var (
	// allEventColumns is the list of fields for the event entity.
	allEventColumns = []string{
		"id",
		"title",
		"created_at",
		"finished_at",
		"description",
		"owner_id",
		"notify_before",
		"is_notified",
	}
)

// GetEvents returns the events by filter.
func (s *PgStorage) GetEvents(_ context.Context, filter eventFilter.Filter) ([]*eventModel.Event, error) {
	sb := squirrel.
		Select(allEventColumns...).
		From(eventsTable).
		PlaceholderFormat(squirrel.Dollar).
		OrderBy("id ASC")

	if filter.Limit > 0 {
		sb = sb.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		sb = sb.Offset(filter.Offset)
	}
	if len(filter.IDs) > 0 {
		sb = sb.Where(squirrel.Eq{"id": filter.IDs})
	}
	if len(filter.OwnerIDs) > 0 {
		sb = sb.Where(squirrel.Eq{"owner_id": filter.OwnerIDs})
	}
	if filter.CreatedFrom != nil {
		sb = sb.Where(squirrel.GtOrEq{"created_at": filter.CreatedFrom})
	}
	if filter.CreatedTo != nil {
		sb = sb.Where(squirrel.Lt{"created_at": filter.CreatedTo})
	}
	if filter.FinishedFrom != nil {
		sb = sb.Where(squirrel.GtOrEq{"finished_at": filter.FinishedFrom})
	}
	if filter.FinishedTo != nil {
		sb = sb.Where(squirrel.Lt{"finished_at": filter.FinishedTo})
	}

	if filter.IsNotified != nil {
		sb = sb.Where(squirrel.Eq{"is_notified": *filter.IsNotified})
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
		err = rows.Scan(
			&event.ID,
			&event.Title,
			&event.CreatedAt,
			&event.FinishedAt,
			&event.Description,
			&event.OwnerID,
			&event.NotifyBefore,
			&event.IsNotified,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	return events, nil
}
