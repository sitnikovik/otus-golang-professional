//go:build integration
// +build integration

package pgsql

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
	timeUtils "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/utils/time"
)

func TestIntegrationPgStorage_UpdateEvent(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		// Create event.
		s := newTestPgStorage(t)
		now := time.Now()
		finishedAt := timeUtils.EndOfWeek(now)
		event := &event.Event{
			Title:      "test",
			CreatedAt:  time.Now(),
			FinishedAt: &finishedAt,
			OwnerID:    1,
		}

		id, err := s.CreateEvent(context.Background(), event)
		require.NoErrorf(t, err, "failed to create event: %v", err)
		require.NotZerof(t, id, "event id is zero")

		// Update event.
		event.ID = id
		event.Title = "test updated"
		err = s.UpdateEvent(ctx, event)
		require.NoErrorf(t, err, "failed to delete event: %v", err)

		// Get event.
		gotEvent, err := s.GetEvent(ctx, id)
		require.NoErrorf(t, err, "failed to get event: %v", err)
		require.Equalf(t, event.Title, gotEvent.Title, "events are not equal")
	})
}
