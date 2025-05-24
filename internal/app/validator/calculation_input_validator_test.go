package validator

import (
	"testing"

	"github.com/KasperSaaby/calculatron-service/internal/app/models"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/stretchr/testify/assert"
)

func Test_CalculationInputValidator_Validate(t *testing.T) {
	testCases := []struct {
		name          string
		operationType string
		operands      []float64
		precision     int
		expectErr     bool
		errMsg        string
	}{
		{
			name:          "empty operands",
			operationType: values.OperationType_Add.String(),
			operands:      []float64{},
			precision:     2,
			expectErr:     true,
			errMsg:        "no operands provided",
		},
		{
			name:          "negative precision",
			operationType: values.OperationType_Add.String(),
			operands:      []float64{1, 2},
			precision:     -1,
			expectErr:     true,
			errMsg:        "precision cannot be negative",
		},
		{
			name:          "valid input",
			operationType: values.OperationType_Add.String(),
			operands:      []float64{1, 2},
			precision:     2,
			expectErr:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := NewCalculationInputValidator()
			input := models.NewCalculationInput(tc.operationType, tc.operands, tc.precision)

			err := validator.Validate(input)

			if tc.expectErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
				return
			}

			assert.NoError(t, err)
		})
	}
}
