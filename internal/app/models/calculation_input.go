package models

import "github.com/KasperSaaby/calculatron-service/internal/domain/values"

type CalculationInput struct {
	operationType values.OperationType
	operands      []float64
	precision     int
}

func NewCalculationInput(operationType string, operands []float64, precision int) CalculationInput {
	return CalculationInput{
		operationType: values.OperationType(operationType),
		operands:      operands,
		precision:     precision,
	}
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
