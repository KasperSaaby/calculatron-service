package validate

import "fmt"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func NumberGreaterThan[V Number](value V, greaterThan V) ValidatorFunc {
	return NumberGreaterThanf(value, greaterThan, "value of %T is not greater than %v", value, greaterThan)
}

func NumberGreaterThanf[V Number](value V, greaterThan V, format string, args ...any) ValidatorFunc {
	return func() error {
		if value > greaterThan {
			return nil
		}

		return fmt.Errorf(format, args...)
	}
}
