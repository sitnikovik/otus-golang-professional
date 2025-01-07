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

func TestImplementation_GetEvent(t *testing.T) {
	t.Parallel()

	type fields struct {
		eventServiceMockFunc func(t *testing.T, eventCreatedAt time.Time) eventService
	}
	type args struct {
		ctx context.Context
		req *pkg.GetEventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pkg.GetEventResponse
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				eventServiceMockFunc: func(t *testing.T, eventCreatedAt time.Time) eventService {
					es := NewMockeventService(t)

					es.EXPECT().
						GetEvent(mock.Anything, mock.Anything).
						Return(&event.Event{
							ID:          1,
							Title:       "title",
							Description: "description",
							CreatedAt:   eventCreatedAt,
							OwnerID:     1,
						}, nil).
						Once()

					return es
				},
			},
			args: args{
				ctx: context.Background(),
				req: &pkg.GetEventRequest{
					Id: 1,
				},
			},
			want: &pkg.GetEventResponse{
				Event: &pkg.Event{
					Id:          1,
					Title:       "title",
					Description: "description",
					CreatedAt:   ToGRPCTime(time.Now()),
					OwnerId:     1,
				},
			},
		},
		{
			name: "err on service errored",
			fields: fields{
				eventServiceMockFunc: func(t *testing.T, eventCreatedAt time.Time) eventService {
					es := NewMockeventService(t)

					es.EXPECT().
						GetEvent(mock.Anything, mock.Anything).
						Return(nil, errors.New("service fake error")).
						Once()

					return es
				},
			},
			args: args{
				ctx: context.Background(),
				req: &pkg.GetEventRequest{
					Id: 0,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			i := &Implementation{
				eventService: tt.fields.eventServiceMockFunc(
					t,
					tt.want.GetEvent().GetCreatedAt().AsTime(),
				),
			}
			got, err := i.GetEvent(tt.args.ctx, tt.args.req)

			require.Equalf(t, tt.wantErr, err != nil, "unexpected error: %v", err)
			require.Equal(t, tt.want, got)
		})
	}
}
