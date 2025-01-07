package event

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	eventFilter "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/filter/event"
	eventModel "github.com/sitnikovik/otus-golang-professional/hw12_13_14_15_calendar/internal/model/event"
)

func TestService_GetEvents(t *testing.T) {
	t.Parallel()

	type fields struct {
		dbMockFunc func(t *testing.T) eventDB
	}
	type args struct {
		ctx    context.Context
		filter eventFilter.Filter
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*eventModel.Event
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				dbMockFunc: func(t *testing.T) eventDB {
					db := NewMockeventDB(t)

					db.EXPECT().
						GetEvents(mock.Anything, mock.Anything).
						Return([]*eventModel.Event{
							{
								ID: 1,
							},
							{
								ID: 2,
							},
						}, nil).
						Once()

					return db
				},
			},
			args: args{
				ctx: context.Background(),
				filter: eventFilter.Filter{
					IDs:   []uint64{1, 2},
					Limit: 2,
				},
			},
			want: []*eventModel.Event{
				{
					ID: 1,
				},
				{
					ID: 2,
				},
			},
		},
		{
			name: "err on get events",
			fields: fields{
				dbMockFunc: func(t *testing.T) eventDB {
					db := NewMockeventDB(t)

					db.EXPECT().
						GetEvents(mock.Anything, mock.Anything).
						Return(nil, errors.New("db fake error")).
						Once()

					return db
				},
			},
			args: args{
				ctx: context.Background(),
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
			got, err := s.GetEvents(tt.args.ctx, tt.args.filter)

			require.Equalf(t, tt.wantErr, err != nil, "unexpected error: %v", err)
			require.Equal(t, tt.want, got)
		})
	}
}
