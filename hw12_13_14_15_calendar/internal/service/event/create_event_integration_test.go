//go:build integration
// +build integration

package event

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/config"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/connections/pg"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/storage/event/pgsql"
)

func TestService_CreateEvent(t *testing.T) {
	t.Parallel()

	type fields struct {
		dbMockFunc func(t *testing.T) eventDB
	}
	type args struct {
		ctx   context.Context
		event *eventModel.Event
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantZero bool
		wantErr  bool
	}{
		{
			name: "ok",
			fields: fields{
				dbMockFunc: func(t *testing.T) eventDB {
					cfg, err := config.NewTestConfig()
					if err != nil {
						t.Fatalf("failed to create test config: %v", err)
					}

					pgPort, _ := strconv.Atoi(cfg.PG.Port)
					conn, err := pg.NewConnPool(cfg.PG.Database, cfg.PG.User, cfg.PG.Password, cfg.PG.Host, pgPort)
					if err != nil {
						t.Fatalf("failed to create pg conn pool: %v", err)
					}

					return pgsql.New(conn)
				},
			},
			args: args{
				ctx: context.Background(),
				event: &eventModel.Event{
					Title:   "test",
					OwnerID: 1,
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Service{
				db: tt.fields.dbMockFunc(t),
			}
			got, err := s.CreateEvent(tt.args.ctx, tt.args.event)

			require.Equalf(t, tt.wantErr, err != nil, "unexpected error: %v", err)
			require.NotZerof(t, got, "unexpected result: got \"%d\" but want not zero", got)
		})
	}
}
