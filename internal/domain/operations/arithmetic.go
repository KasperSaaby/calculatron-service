package operations

import (
	"math"

	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
)

func Add(operands ...float64) (float64, error) {
	if len(operands) < 2 {
		return 0, values.NewDomainError("Addition requires at least two operands", values.Code_TwoOrMoreOperandsRequired)
	}
	result := 0.0
	for _, operand := range operands {
		result += operand
	}
	return result, nil
}

func Subtract(operands ...float64) (float64, error) {
	if len(operands) != 2 {
		return 0, values.NewDomainError("Subtraction requires exactly two operands", values.Code_TwoOperandsRequired)
	}
	return operands[0] - operands[1], nil
}

func Multiply(operands ...float64) (float64, error) {
	if len(operands) < 2 {
		return 0, values.NewDomainError("Multiplication requires at least two operands", values.Code_TwoOperandsRequired)
	}
	result := 1.0
	for _, operand := range operands {
		result *= operand
	}
	return result, nil
}

func Divide(operands ...float64) (float64, error) {
	if len(operands) != 2 {
		return 0, values.NewDomainError("Division requires exactly two operands", values.Code_TwoOperandsRequired)
	}
	if operands[1] == 0 {
		return 0, values.NewDomainError("Cannot divide by zero", values.Code_DivisionByZero)
	}
	return operands[0] / operands[1], nil
}

func Power(operands ...float64) (float64, error) {
	if len(operands) != 2 {
		return 0, values.NewDomainError("Power operation requires exactly two operands (base and exponent)", values.Code_TwoOperandsRequired)
	}
	return math.Pow(operands[0], operands[1]), nil
}
