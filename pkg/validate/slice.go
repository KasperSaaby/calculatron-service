package validate

import "fmt"

func SliceLenGreaterThan[T comparable](slice []T, greaterThan int) ValidatorFunc {
	return SliceLenGreaterThanf(slice, greaterThan, "length of slice is not greater than %v", greaterThan)
}

func SliceLenGreaterThanf[T comparable](slice []T, greaterThan int, format string, args ...any) ValidatorFunc {
	return func() error {
		if len(slice) > greaterThan {
			return nil
		}

		return fmt.Errorf(format, args...)
	}
}
