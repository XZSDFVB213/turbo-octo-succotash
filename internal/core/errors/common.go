package core_errors

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrInvalidArgument = errors.New("Invalid argument")
	ErrConflict        = errors.New("Conflict")
)
