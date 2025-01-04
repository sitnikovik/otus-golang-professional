package grpc

import (
	"context"
	"errors"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

func TestImplementation_CreateEvent(t *testing.T) {
	t.Parallel()

	type fields struct {
		eventServiceMockFunc func(t *testing.T) eventService
	}
	type args struct {
		ctx context.Context
		req *pkg.CreateEventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pkg.CreateEventResponse
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				eventServiceMockFunc: func(t *testing.T) eventService {
					es := NewMockeventService(t)

					es.EXPECT().
						CreateEvent(mock.Anything, mock.Anything).
						Return(uint64(123), nil).
						Once()

					return es
				},
			},
			args: args{
				ctx: context.Background(),
				req: &pkg.CreateEventRequest{
					Title:       "title",
					Description: "description",
					OwnerId:     1,
				},
			},
			want: &pkg.CreateEventResponse{
				Id: 123,
			},
		},
		{
			name: "err on service errored",
			fields: fields{
				eventServiceMockFunc: func(t *testing.T) eventService {
					es := NewMockeventService(t)

					es.EXPECT().
						CreateEvent(mock.Anything, mock.Anything).
						Return(uint64(0), errors.New("service fake error")).
						Once()

					return es
				},
			},
			args: args{
				ctx: context.Background(),
				req: &pkg.CreateEventRequest{
					Title:       "title",
					Description: "description",
					OwnerId:     1,
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
				eventService: tt.fields.eventServiceMockFunc(t),
			}
			got, err := i.CreateEvent(tt.args.ctx, tt.args.req)

			require.Equalf(t, tt.wantErr, err != nil, "unexpected error: %v", err)
			require.Equal(t, tt.want, got)
		})
	}
}
