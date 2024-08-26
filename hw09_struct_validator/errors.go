package hw09structvalidator

import (
	"errors"
)

var (
	ErrInvalidLength  = errors.New("invalid length")
	ErrNotMatchRegexp = errors.New("not match regexp")
	ErrNotInRange     = errors.New("not in")
	ErrNotGreater     = errors.New("not greater")
	ErrNotLesser      = errors.New("not lesser")
)
