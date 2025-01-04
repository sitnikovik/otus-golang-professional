package grpc

import (
	"context"
	"errors"
	"testing"

	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	pkg "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/pkg/calendar/v1"
)

func TestImplementation_UpdateEvent(t *testing.T) {
	t.Parallel()

	type fields struct {
		eventServiceMockFunc func(t *testing.T) eventService
	}
	type args struct {
		ctx context.Context
		req *pkg.UpdateEventRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pkg.UpdateEventResponse
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				eventServiceMockFunc: func(t *testing.T) eventService {
					es := NewMockeventService(t)

					es.EXPECT().
						UpdateEvent(mock.Anything, mock.Anything).
						Return(nil).
						Once()

					return es
				},
			},
			args: args{
				ctx: context.Background(),
				req: &pkg.UpdateEventRequest{
					Event: &pkg.Event{
						Id:          1,
						Title:       "title",
						Description: "description",
						OwnerId:     1,
					},
				},
			},
			want: &pkg.UpdateEventResponse{},
		},
		{
			name: "err on service errored",
			fields: fields{
				eventServiceMockFunc: func(t *testing.T) eventService {
					es := NewMockeventService(t)

					es.EXPECT().
						UpdateEvent(mock.Anything, mock.Anything).
						Return(errors.New("service fake error")).
						Once()

					return es
				},
			},
			args: args{
				ctx: context.Background(),
				req: &pkg.UpdateEventRequest{
					Event: &pkg.Event{
						Id:          1,
						Title:       "title",
						Description: "description",
						OwnerId:     1,
					},
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
			got, err := i.UpdateEvent(tt.args.ctx, tt.args.req)

			require.Equalf(t, tt.wantErr, err != nil, "unexpected error: %v", err)
			require.Equal(t, tt.want, got)
		})
	}
}
