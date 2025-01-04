package event

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

func TestService_GetEvent(t *testing.T) {
	t.Parallel()

	type fields struct {
		dbMockFunc func(t *testing.T) eventDB
	}
	type args struct {
		ctx     context.Context
		eventID uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *eventModel.Event
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				dbMockFunc: func(t *testing.T) eventDB {
					db := NewMockeventDB(t)

					db.EXPECT().
						GetEvent(mock.Anything, mock.Anything).
						Return(&eventModel.Event{
							ID: 1,
						}, nil).
						Once()

					return db
				},
			},
			args: args{
				ctx:     context.Background(),
				eventID: 1,
			},
			want: &eventModel.Event{
				ID: 1,
			},
		},
		{
			name: "err on get event",
			fields: fields{
				dbMockFunc: func(t *testing.T) eventDB {
					db := NewMockeventDB(t)

					db.EXPECT().
						GetEvent(mock.Anything, mock.Anything).
						Return(&eventModel.Event{
							ID: 1,
						}, nil).
						Once()

					return db
				},
			},
			args: args{
				ctx:     context.Background(),
				eventID: 1,
			},
			want: &eventModel.Event{
				ID: 1,
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
			got, err := s.GetEvent(tt.args.ctx, tt.args.eventID)

			require.Equalf(t, tt.wantErr, err != nil, "unexpected error: %v", err)
			require.Equalf(t, tt.want, got, "unexpected result: got %v, want %v", got, tt.want)
		})
	}
}
