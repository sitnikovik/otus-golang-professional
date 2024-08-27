package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name        string
		in          interface{}
		expectedErr error
	}{
		{
			name: "valid user",
			in: User{
				ID:    "123456789012345678901234567890123456",
				Email: "test@test.ru",
				Age:   50,
				Role:  "admin",
				Phones: []string{
					"12312312312",
				},
			},
			expectedErr: nil,
		},
		{
			name: "invalid user",
			in: User{
				ID:    "123456789012345678901",
				Email: "@test.ru",
				Age:   17,
				Role:  "user",
				Phones: []string{
					"12312312312",
					"12312",
				},
			},
			expectedErr: ValidationErrors{
				{Field: "ID", Err: ErrInvalidLength},
				{Field: "Age", Err: ErrNotGreater},
				{Field: "Email", Err: ErrNotMatchRegexp},
				{Field: "Role", Err: ErrNotInRange},
				{Field: "Phones", Err: fmt.Errorf("elem 1: %v", ErrInvalidLength)},
			},
		},
		{
			name: "valid app",
			in: App{
				Version: "1.0.0",
			},
			expectedErr: nil,
		},
		{
			name: "invalid app",
			in: App{
				Version: "1.0",
			},
			expectedErr: ValidationErrors{
				{Field: "Version", Err: ErrInvalidLength},
			},
		},
		{
			name: "valid token",
			in: Token{
				Header:    []byte("header"),
				Payload:   []byte("payload"),
				Signature: []byte("signature"),
			},
			expectedErr: nil,
		},
		{
			name: "valid response",
			in: Response{
				Code: 200,
				Body: "body",
			},
			expectedErr: nil,
		},
		{
			name: "invalid response",
			in: Response{
				Code: 201,
				Body: "body",
			},
			expectedErr: ValidationErrors{
				{Field: "Code", Err: ErrNotInRange},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()

			require.Equal(t, tt.expectedErr, Validate(tt.in))
		})
	}
}
