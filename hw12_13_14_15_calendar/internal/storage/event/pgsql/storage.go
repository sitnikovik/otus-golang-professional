package pgsql

import (
	"context"

	"github.com/jackc/pgx"
)

const (
	// eventsTable is the name of the table for events.
	eventsTable = "events"
)

// PgStorage is the storage for the events.
type PgStorage struct {
	// db is the connection pool to the database.
	db *pgx.ConnPool
}

// New creates and returns the sql storage instance.
func New(pg *pgx.ConnPool) *PgStorage {
	return &PgStorage{
		db: pg,
	}
}

// Close closes the storage.
func (s *PgStorage) Close(_ context.Context) error {
	s.db.Close()
	return nil
}
