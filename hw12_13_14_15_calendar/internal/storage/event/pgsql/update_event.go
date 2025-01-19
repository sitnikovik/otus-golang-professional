package pgsql

import (
	"context"

	"github.com/Masterminds/squirrel"

	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

// UpdateEvent updates the event.
func (s *PgStorage) UpdateEvent(_ context.Context, event *eventModel.Event) error {
	sb := squirrel.
		Update(eventsTable).
		PlaceholderFormat(squirrel.Dollar).
		Set("title", event.Title).
		Set("description", event.Description).
		Set("finished_at", event.FinishedAt).
		Set("notify_before", event.NotifyBefore).
		Set("is_notified", event.IsNotified).
		Where(squirrel.Eq{"id": event.ID})

	sql, args, err := sb.ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}
