//go:build integration
// +build integration

package pgsql

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	eventFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
	timeUtils "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/utils/time"
)

func TestPgStorage_GetEvents(t *testing.T) {
	t.Parallel()

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		filter := eventFilter.Filter{
			Limit: 3,
		}
		filter.IDs = make([]uint64, 0, filter.Limit)

		s := newTestPgStorage(t)

		// Create events.
		now := time.Now()
		n := int(filter.Limit)
		for i := 0; i < n; i++ {
			finishedAt := timeUtils.EndOfWeek(now)
			id, err := s.CreateEvent(context.Background(), &event.Event{
				Title:      "test",
				CreatedAt:  time.Now(),
				FinishedAt: &finishedAt,
				OwnerID:    1,
			})

			require.NoErrorf(t, err, "failed to create event: %v", err)
			require.NotZerof(t, id, "event id is zero")

			filter.IDs = append(filter.IDs, id)
		}

		// Get events.
		events, err := s.GetEvents(ctx, filter)

		require.NoErrorf(t, err, "failed to get events: %v", err)
		require.Equalf(t, n, len(events), "events count is not equal")
	})
}
