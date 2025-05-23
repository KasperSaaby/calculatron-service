package validator

import (
	"github.com/KasperSaaby/calculatron-service/internal/app/models"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
)

type CalculationInputValidator struct{}

func NewCalculationInputValidator() *CalculationInputValidator {
	return &CalculationInputValidator{}
}

func (v *CalculationInputValidator) Validate(input models.CalculationInput) error {
	if len(input.Operands()) == 0 {
		return NewValidationError("no operands provided")
	}

	if input.Precision() < 0 {
		return NewValidationError("precision cannot be negative")
	}

	// Validate operation type
	switch input.OperationType() {
	case
		values.OperationType_Add,
		values.OperationType_Subtract,
		values.OperationType_Multiply,
		values.OperationType_Divide,
		values.OperationType_Power:
		// Valid operation types
	default:
		return NewValidationError("unsupported operation type: %s", input.OperationType())
	}

	return nil
}
