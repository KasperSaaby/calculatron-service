package models

import (
	"fmt"

	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
)

type CalculationInput struct {
	operationType values.OperationType
	operands      []float64
	precision     int
}

func NewCalculationInput(operationType string, operands []float64, precision int) (CalculationInput, error) {
	input := CalculationInput{
		operationType: values.OperationType(operationType),
		operands:      operands,
		precision:     precision,
	}

	if err := input.Validate(); err != nil {
		return CalculationInput{}, err
	}

	return input, nil
}

func (i CalculationInput) Validate() error {
	if len(i.operands) == 0 {
		return fmt.Errorf("no operands provided")
	}

	if i.precision < 0 {
		return fmt.Errorf("precision cannot be negative")
	}

	// Validate operation type
	switch i.operationType {
	case
		values.OperationType_Add,
		values.OperationType_Subtract,
		values.OperationType_Multiply,
		values.OperationType_Divide,
		values.OperationType_Power:
		// Valid operation types
	default:
		return fmt.Errorf("unsupported operation type: %s", i.operationType)
	}

	return nil
}

func (i CalculationInput) OperationType() values.OperationType {
	return i.operationType
}

func (i CalculationInput) Operands() []float64 {
	return i.operands
}

func (i CalculationInput) Precision() int {
	return i.precision
}
