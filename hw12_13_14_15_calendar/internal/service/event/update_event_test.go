package event

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

func TestService_UpdateEvent(t *testing.T) {
	t.Parallel()

	type fields struct {
		dbMockFunc func(t *testing.T) eventDB
	}
	type args struct {
		ctx   context.Context
		event *eventModel.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				dbMockFunc: func(t *testing.T) eventDB {
					db := NewMockeventDB(t)

					db.EXPECT().
						UpdateEvent(mock.Anything, mock.Anything).
						Return(nil).
						Once()

					return db
				},
			},
			args: args{
				ctx: context.Background(),
				event: &eventModel.Event{
					ID: 1,
				},
			},
		},
		{
			name: "err on update event",
			fields: fields{
				dbMockFunc: func(t *testing.T) eventDB {
					db := NewMockeventDB(t)

					db.EXPECT().
						UpdateEvent(mock.Anything, mock.Anything).
						Return(errors.New("db fake error")).
						Once()

					return db
				},
			},
			args: args{
				ctx: context.Background(),
				event: &eventModel.Event{
					ID: 1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Service{
				db: tt.fields.dbMockFunc(t),
			}
			err := s.UpdateEvent(tt.args.ctx, tt.args.event)

			require.Equalf(t, tt.wantErr, err != nil, "unexpected error: %v", err)
		})
	}
}
