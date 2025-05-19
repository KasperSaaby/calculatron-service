package operations

import (
	"fmt"
	"math"
)

// Add performs addition of two or more numbers.
func Add(operands ...float64) (float64, error) {
	if len(operands) < 2 {
		return 0, fmt.Errorf("addition requires at least two operands")
	}
	result := 0.0
	for _, operand := range operands {
		result += operand
	}
	return result, nil
}

// Subtract performs subtraction. It expects exactly two operands.
func Subtract(operands ...float64) (float64, error) {
	if len(operands) != 2 {
		return 0, fmt.Errorf("subtraction requires exactly two operands")
	}
	return operands[0] - operands[1], nil
}

// Multiply performs multiplication of two or more numbers.
func Multiply(operands ...float64) (float64, error) {
	if len(operands) < 2 {
		return 0, fmt.Errorf("multiplication requires at least two operands")
	}
	result := 1.0
	for _, operand := range operands {
		result *= operand
	}
	return result, nil
}

// Divide performs division. It expects exactly two operands, and the divisor cannot be zero.
func Divide(operands ...float64) (float64, error) {
	if len(operands) != 2 {
		return 0, fmt.Errorf("division requires exactly two operands")
	}
	if operands[1] == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return operands[0] / operands[1], nil
}

// Power calculates base raised to the power of exponent. It expects exactly two operands.
func Power(operands ...float64) (float64, error) {
	if len(operands) != 2 {
		return 0, fmt.Errorf("power operation requires exactly two operands (base and exponent)")
	}
	return math.Pow(operands[0], operands[1]), nil
}
