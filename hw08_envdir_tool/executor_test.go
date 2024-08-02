package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	type args struct {
		cmd []string
		env Environment
	}
	tests := []struct {
		name     string
		args     args
		wantCode int
	}{
		{
			name: "ok on echo the testdata dir",
			args: args{
				cmd: []string{"/bin/bash", "./testdata/echo.sh"},
				env: Environment{},
			},
		},
		{
			name: "ok on testsdata envdir specified but no cmd",
			args: args{
				env: Environment{},
			},
		},
		{
			name: "err on nil env",
			args: args{
				cmd: []string{"/bin/bash", "./testdata/echo.sh"},
				env: nil,
			},
			wantCode: ErrCode,
		},
		{
			name: "err on unknown cmd",
			args: args{
				cmd: []string{"1213"},
				env: Environment{},
			},
			wantCode: ErrCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.wantCode, RunCmd(tt.args.cmd, tt.args.env))
		})
	}
}
