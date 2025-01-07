package event

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

func TestService_CreateEvent(t *testing.T) {
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
						CreateEvent(mock.Anything, mock.Anything).
						Return(uint64(1), nil)

					return db
				},
			},
			args: args{
				ctx: context.Background(),
				event: &eventModel.Event{
					Title:   "test",
					OwnerID: 1,
				},
			},
		},
		{
			name: "err on nil event",
			fields: fields{
				dbMockFunc: func(t *testing.T) eventDB {
					return NewMockeventDB(t)
				},
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: true,
		},
		{
			name: "validation error",
			fields: fields{
				dbMockFunc: func(t *testing.T) eventDB {
					return NewMockeventDB(t)
				},
			},
			args: args{
				ctx: context.Background(),
				event: &eventModel.Event{
					Title: "test",
				},
			},
			wantErr: true,
		},
		{
			name: "err on db create event",
			fields: fields{
				dbMockFunc: func(t *testing.T) eventDB {
					db := NewMockeventDB(t)

					db.EXPECT().
						CreateEvent(mock.Anything, mock.Anything).
						Return(uint64(0), errors.New("db fake error"))

					return db
				},
			},
			args: args{
				ctx: context.Background(),
				event: &eventModel.Event{
					Title:      "test",
					CreatedAt:  time.Now(),
					FinishedAt: timeNowWithPointer(),
					OwnerID:    1,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			s := &Service{
				db: tt.fields.dbMockFunc(t),
			}
			got, err := s.CreateEvent(tt.args.ctx, tt.args.event)

			require.Equalf(t, tt.wantErr, err != nil, "unexpected error: %v", err)
			if tt.wantErr {
				require.Zerof(t, got, "expected zero result but got \"%d\"", got)
				return
			}
			require.NotZerof(t, got, "unexpected result: got \"%d\" but want not zero", got)
		})
	}
}

func Test_validateEventToCreate(t *testing.T) {
	t.Parallel()

	type args struct {
		event *eventModel.Event
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				event: &eventModel.Event{
					Title:      "test",
					CreatedAt:  time.Now(),
					FinishedAt: timeNowWithPointer(),
					OwnerID:    1,
				},
			},
			wantErr: false,
		},
		{
			name: "nil event",
			args: args{
				event: nil,
			},
			wantErr: true,
		},
		{
			name: "empty title",
			args: args{
				event: &eventModel.Event{
					Title: "",
				},
			},
			wantErr: true,
		},
		{
			name: "zero created date at",
			args: args{
				event: &eventModel.Event{
					Title: "test",
				},
			},
			wantErr: true,
		},
		{
			name: "zero finished date at",
			args: args{
				event: &eventModel.Event{
					Title:      "test",
					CreatedAt:  time.Now(),
					FinishedAt: &time.Time{},
				},
			},
			wantErr: true,
		},
		{
			name: "zero owner id",
			args: args{
				event: &eventModel.Event{
					Title:      "test",
					CreatedAt:  time.Now(),
					FinishedAt: timeNowWithPointer(),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := validateEventToCreate(tt.args.event)

			require.Equalf(t, tt.wantErr, err != nil, "unexpected error: %v", err)
		})
	}
}

// timeNowWithPointer returns the current time with a pointer.
func timeNowWithPointer() *time.Time {
	now := time.Now()
	return &now
}
