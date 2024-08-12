package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args

		envPreset []string // Envnames to be set before targer func done to check NeedRemove flag for env vars parsed

		want    Environment
		wantErr bool
	}{
		{
			name: "err on filepath specified",
			args: args{
				path: "testdata/env/BAR",
			},
			wantErr: true,
		},
		{
			name: "err on unexisting dir specified",
			args: args{
				path: "testdata/envdir",
			},
			wantErr: true,
		},
		{
			name: "ok on testdata envdir specified",
			args: args{
				path: "testdata/env",
			},
			want: Environment{
				"BAR": EnvValue{
					Value: "bar",
				},
				"EMPTY": EnvValue{
					Value: "",
				},
				"FOO": EnvValue{
					Value: "   foo\nwith new line",
				},
				"HELLO": EnvValue{
					Value: "\"hello\"",
				},
				"UNSET": EnvValue{
					Value: "",
				},
			},
		},
		{
			name: "ok on testdata envdir specified with env vars preset before",
			args: args{
				path: "testdata/env",
			},
			envPreset: []string{"BAR", "UNSET"},
			want: Environment{
				"BAR": EnvValue{
					Value:      "bar",
					NeedRemove: true,
				},
				"EMPTY": EnvValue{
					Value: "",
				},
				"FOO": EnvValue{
					Value: "   foo\nwith new line",
				},
				"HELLO": EnvValue{
					Value: "\"hello\"",
				},
				"UNSET": EnvValue{
					Value:      "",
					NeedRemove: true,
				},
			},
		},
	}

	for i := 0; i < len(tests); i++ {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			for _, v := range tt.envPreset {
				os.Setenv(v, "")
			}

			got, err := ReadDir(tt.args.path)

			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, got)
				return
			}

			for envName, envValue := range got {
				if wantValue, ok := tt.want[envName]; ok {
					require.Equal(t, wantValue, envValue)
				}
			}
		})
	}
}
