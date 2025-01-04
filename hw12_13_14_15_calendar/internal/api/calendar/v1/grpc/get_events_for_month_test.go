package grpc

import (
	"context"
	"errors"
	"testing"
	"time"

	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

func TestImplementation_GetEventsForMonth(t *testing.T) {
	t.Parallel()

	now := time.Now()
	type fields struct {
		eventServiceMockFunc func(t *testing.T) eventService
	}
	type args struct {
		ctx context.Context
		req *pkg.GetEventsForMonthRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pkg.GetEventsResponse
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				eventServiceMockFunc: func(t *testing.T) eventService {
					s := NewMockeventService(t)

					s.EXPECT().
						GetEventsForMonth(mock.Anything).
						Return([]*event.Event{
							{
								ID:        1,
								CreatedAt: now,
							},
							{
								ID:        2,
								CreatedAt: now,
							},
						}, nil).
						Once()

					return s
				},
			},
			args: args{
				ctx: context.Background(),
				req: &pkg.GetEventsForMonthRequest{},
			},
			want: &pkg.GetEventsResponse{
				Events: []*pkg.Event{
					{
						Id:        1,
						CreatedAt: ToGRPCTime(now),
					},
					{
						Id:        2,
						CreatedAt: ToGRPCTime(now),
					},
				},
				Total: 2,
			},
		},
		{
			name: "err service errored",
			fields: fields{
				eventServiceMockFunc: func(t *testing.T) eventService {
					s := NewMockeventService(t)

					s.EXPECT().
						GetEventsForMonth(mock.Anything).
						Return(nil, errors.New("service fake error")).
						Once()

					return s
				},
			},
			args: args{
				ctx: context.Background(),
				req: &pkg.GetEventsForMonthRequest{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Implementation{
				eventService: tt.fields.eventServiceMockFunc(t),
			}
			got, err := i.GetEventsForMonth(tt.args.ctx, tt.args.req)

			require.Equalf(t, tt.wantErr, err != nil, "unexpected error: %v", err)
			require.Equal(t, tt.want, got)
		})
	}
}
