package event

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestService_DeleteEvent(t *testing.T) {
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
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				dbMockFunc: func(t *testing.T) eventDB {
					db := NewMockeventDB(t)

					db.EXPECT().
						DeleteEvent(mock.Anything, mock.Anything).
						Return(nil).
						Once()

					return db
				},
			},
		},
		{
			name: "err on delete event",
			fields: fields{
				dbMockFunc: func(t *testing.T) eventDB {
					db := NewMockeventDB(t)

					db.EXPECT().
						DeleteEvent(mock.Anything, mock.Anything).
						Return(errors.New("db fake error")).
						Once()

					return db
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
			err := s.DeleteEvent(tt.args.ctx, tt.args.eventID)

			require.Equalf(t, tt.wantErr, err != nil, "unexpected error: %v", err)
		})
	}
}
