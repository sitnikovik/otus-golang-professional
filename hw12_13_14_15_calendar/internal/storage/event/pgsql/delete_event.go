package pgsql

import (
	"context"

	"github.com/Masterminds/squirrel"
)

// DeleteEvent deletes the event.
func (s *PgStorage) DeleteEvent(_ context.Context, eventID uint64) error {
	sb := squirrel.
		Delete(eventsTable).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{"id": eventID})

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
