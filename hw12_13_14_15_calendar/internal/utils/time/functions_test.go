package time

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestEndOfWeek(t *testing.T) {
	t.Parallel()

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "monday -> sunday",
			args: args{
				t: time.Date(2021, 1, 4, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2021, 1, 10, 23, 59, 59, 0, time.UTC),
		},
		{
			name: "sunday -> sunday",
			args: args{
				t: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2021, 1, 3, 23, 59, 59, 0, time.UTC),
		},
		{
			name: "same sundays",
			args: args{
				t: time.Date(2021, 1, 3, 23, 59, 59, 0, time.UTC),
			},
			want: time.Date(2021, 1, 3, 23, 59, 59, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := EndOfWeek(tt.args.t)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestEndOfDay(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "17:00:00 -> 23:59:59",
			args: args{
				t: time.Date(2021, 1, 1, 17, 0, 0, 0, time.UTC),
			},
			want: time.Date(2021, 1, 1, 23, 59, 59, 0, time.UTC),
		},
		{
			name: "23:59:59 -> 23:59:59",
			args: args{
				t: time.Date(2021, 1, 1, 23, 59, 59, 0, time.UTC),
			},
			want: time.Date(2021, 1, 1, 23, 59, 59, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := EndOfDay(tt.args.t)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestEndOfMonth(t *testing.T) {
	t.Parallel()

	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "2021-01-15 -> 2021-01-31 23:59:59",
			args: args{
				t: time.Date(2021, 1, 15, 0, 0, 0, 0, time.UTC),
			},
			want: time.Date(2021, 1, 31, 23, 59, 59, 0, time.UTC),
		},
		{
			name: "2021-01-31 23:59:59 -> 2021-01-31 23:59:59",
			args: args{
				t: time.Date(2021, 1, 31, 23, 59, 59, 0, time.UTC),
			},
			want: time.Date(2021, 1, 31, 23, 59, 59, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := EndOfMonth(tt.args.t)

			require.Equal(t, tt.want, got)
		})
	}
}

func TestUpToDays(t *testing.T) {
	t.Parallel()

	type args struct {
		t    time.Time
		days int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "2021-01-01 + 5 days -> 2021-01-06 00:00:00",
			args: args{
				t:    time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				days: 5,
			},
			want: time.Date(2021, 1, 6, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := UpToDays(tt.args.t, tt.args.days)

			require.Equal(t, tt.want, got)
		})
	}
}
