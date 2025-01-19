package event

import (
	"testing"
	"time"
)

func TestEvent_IsToNotify(t *testing.T) {
	tests := []struct {
		name   string
		fields Event
		want   bool
	}{
		{
			name: "event is already notified",
			fields: Event{
				IsNotified: true,
				FinishedAt: func() *time.Time {
					t := time.Now().Add(10 * time.Minute)
					return &t
				}(),
				NotifyBefore: 5 * time.Minute,
			},
			want: false,
		},
		{
			name: "event is within the notification window",
			fields: Event{
				IsNotified: false,
				FinishedAt: func() *time.Time {
					t := time.Now().Add(4 * time.Minute)
					return &t
				}(),
				NotifyBefore: 5 * time.Minute,
			},
			want: true,
		},
		{
			name: "event is not within the notification window",
			fields: Event{
				IsNotified: false,
				FinishedAt: func() *time.Time {
					t := time.Now().Add(10 * time.Minute)
					return &t
				}(),
				NotifyBefore: 5 * time.Minute,
			},
			want: false,
		},
		{
			name: "event has passed and not notified",
			fields: Event{
				IsNotified: false,
				FinishedAt: func() *time.Time {
					t := time.Now().Add(-10 * time.Minute)
					return &t
				}(),
				NotifyBefore: 5 * time.Minute,
			},
			want: true,
		},
		{
			name: "event has passed and already notified",
			fields: Event{
				IsNotified: true,
				FinishedAt: func() *time.Time {
					t := time.Now().Add(-10 * time.Minute)
					return &t
				}(),
				NotifyBefore: 5 * time.Minute,
			},
			want: false,
		},
		{
			name: "event is exactly at the notification time",
			fields: Event{
				IsNotified: false,
				FinishedAt: func() *time.Time {
					t := time.Now().Add(5 * time.Minute)
					return &t
				}(),
				NotifyBefore: 5 * time.Minute,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Event{
				ID:           tt.fields.ID,
				Title:        tt.fields.Title,
				CreatedAt:    tt.fields.CreatedAt,
				FinishedAt:   tt.fields.FinishedAt,
				Description:  tt.fields.Description,
				OwnerID:      tt.fields.OwnerID,
				NotifyBefore: tt.fields.NotifyBefore,
				IsNotified:   tt.fields.IsNotified,
			}
			if got := e.IsToNotify(); got != tt.want {
				t.Errorf("Event.IsToNotify() = %v, want %v", got, tt.want)
			}
		})
	}
}
