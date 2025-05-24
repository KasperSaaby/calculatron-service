package validator

import "github.com/KasperSaaby/calculatron-service/internal/app/models"

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

	return nil
}
