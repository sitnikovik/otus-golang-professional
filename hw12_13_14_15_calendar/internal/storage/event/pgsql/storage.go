package pgsql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

const (
	// eventsTable is the name of the table for events
	eventsTable = "events"
)

type PgStorage struct {
	db *pgx.Conn // Пул коннектов к БД
}

// New creates and returns the sql storage instance
func New(pg *pgx.Conn) *PgStorage {
	return &PgStorage{
		db: pg,
	}
}

// Close closes the storage
func (s *PgStorage) Close(ctx context.Context) error {
	return s.db.Close(ctx)
}
