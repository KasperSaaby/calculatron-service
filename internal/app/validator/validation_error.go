package validator

import "fmt"

func NewValidationError(format string, args ...any) *ValidationError {
	return &ValidationError{
		format: format,
		args:   args,
	}
}

type ValidationError struct {
	format string
	args   []any
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf(e.format, e.args...)
}
