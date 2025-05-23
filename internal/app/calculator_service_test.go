package app

import (
	"context"
	"testing"

	"github.com/KasperSaaby/calculatron-service/internal/app/models"
	"github.com/KasperSaaby/calculatron-service/internal/app/validator"
	"github.com/KasperSaaby/calculatron-service/internal/domain/operations"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/stretchr/testify/assert"
)

func Test_CalculatorService_PerformCalculation(t *testing.T) {
	testCases := []struct {
		name           string
		operationType  values.OperationType
		operands       []float64
		precision      int
		expectErr      bool
		errMsg         string
		expectedResult float64
	}{
		{
			name:          "empty operands",
			operationType: values.OperationType_Add,
			operands:      []float64{},
			precision:     2,
			expectErr:     true,
			errMsg:        "no operands provided",
		},
		{
			name:          "negative precision",
			operationType: values.OperationType_Add,
			operands:      []float64{1, 2},
			precision:     -1,
			expectErr:     true,
			errMsg:        "precision cannot be negative",
		},
		{
			name:          "unsupported operation",
			operationType: "unsupported",
			operands:      []float64{1, 2},
			precision:     2,
			expectErr:     true,
			errMsg:        "unsupported operation type",
		},
		{
			name:           "successful addition with whole numbers",
			operationType:  values.OperationType_Add,
			operands:       []float64{1, 2},
			precision:      2,
			expectErr:      false,
			expectedResult: 3.00,
		},
		{
			name:           "successful addition with decimals",
			operationType:  values.OperationType_Add,
			operands:       []float64{1.234, 2.345},
			precision:      2,
			expectErr:      false,
			expectedResult: 3.58,
		},
		{
			name:           "successful addition with high precision",
			operationType:  values.OperationType_Add,
			operands:       []float64{1.2345, 2.3456},
			precision:      4,
			expectErr:      false,
			expectedResult: 3.5801,
		},
		{
			name:           "successful addition with multiple operands",
			operationType:  values.OperationType_Add,
			operands:       []float64{1, 2, 3, 4},
			precision:      2,
			expectErr:      false,
			expectedResult: 10.00,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			operationFactory := operations.NewOperationFactory()
			calculationInputValidator := validator.NewCalculationInputValidator()
			sut := NewCalculatorService(operationFactory, calculationInputValidator)

			input := models.NewCalculationInput(tc.operationType.String(), tc.operands, tc.precision)
			result, err := sut.PerformCalculation(ctx, input)

			if tc.expectErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.errMsg)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedResult, result.Result)
			assert.Equal(t, tc.precision, result.Precision)
			assert.NotEmpty(t, result.OperationID)
			assert.NotZero(t, result.Timestamp)
		})
	}
}
