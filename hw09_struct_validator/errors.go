package hw09structvalidator

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrInvalidLength is an error that is returned when the value has an invalid length.
	ErrInvalidLength = errors.New("invalid length")
	// ErrNotMatchRegexp is an error that is returned when the value does not match the specified regexp.
	ErrNotMatchRegexp = errors.New("not match regexp")
	// ErrNotInRange is an error that is returned when the value is not in the specified range.
	ErrNotInRange = errors.New("not in")
	// ErrNotGreater is an error that is returned when the value is not greater than the specified value.
	ErrNotGreater = errors.New("not greater")
	// ErrNotLesser is an error that is returned when the value is not lesser than the specified value.
	ErrNotLesser = errors.New("not lesser")
)

// ValidationError describes a validation error.
type ValidationError struct {
	Field string
	Err   error
}

// ValidationErrors is a list of validation errors.
type ValidationErrors []ValidationError

// Error returns the error message.
func (v ValidationErrors) Error() string {
	ss := make([]string, 0, len(v))
	for _, e := range v {
		ss = append(ss, fmt.Sprintf("%s: %s", e.Field, e.Err.Error()))
	}

	return strings.Join(ss, ", ")
}
