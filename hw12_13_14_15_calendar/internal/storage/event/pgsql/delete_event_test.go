package pgsql

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
	timeUtils "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/utils/time"
)

func TestPgStorage_DeleteEvent(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		// Create event.
		s := newTestPgStorage(t)
		now := time.Now()
		finishedAt := timeUtils.EndOfWeek(now)
		id, err := s.CreateEvent(context.Background(), &event.Event{
			Title:      "test",
			CreatedAt:  time.Now(),
			FinishedAt: &finishedAt,
			OwnerID:    1,
		})
		require.NoErrorf(t, err, "failed to create event: %v", err)
		require.NotZerof(t, id, "event id is zero")

		// Delete event.
		err = s.DeleteEvent(ctx, id)
		require.NoErrorf(t, err, "failed to delete event: %v", err)
	})
}
