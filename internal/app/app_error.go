package app

import "fmt"

func newAppError(format string, args ...any) *AppError {
	return &AppError{
		format: format,
		args:   args,
	}
}

type AppError struct {
	format string
	args   []any
}

func (e *AppError) Error() string {
	return fmt.Sprintf(e.format, e.args...)
}
