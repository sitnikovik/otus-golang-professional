//go:build integration
// +build integration

package pgsql

import (
	"strconv"
	"testing"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/connections/pg"
)

// newTestPgStorage creates new PgStorage for integration testing
func newTestPgStorage(t *testing.T) *PgStorage {
	cfg, err := config.NewTestConfig()
	if err != nil {
		t.Fatalf("failed to create test config: %v", err)
	}

	pgPort, _ := strconv.Atoi(cfg.PG.Port)
	conn, err := pg.NewConnPool(cfg.PG.Database, cfg.PG.User, cfg.PG.Password, cfg.PG.Host, pgPort)
	if err != nil {
		t.Fatalf("failed to create pg conn pool: %v", err)
	}

	return New(conn)
}
