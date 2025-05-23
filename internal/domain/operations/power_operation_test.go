package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_PowerOperation(t *testing.T) {
	testCases := []struct {
		name           string
		operands       []float64
		precision      int
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "successful power operation with positive integers",
			operands:       []float64{2, 3},
			precision:      2,
			expectedResult: 8,
			expectError:    false,
		},
		{
			name:           "successful power operation with decimal exponent",
			operands:       []float64{4, 0.5},
			precision:      2,
			expectedResult: 2,
			expectError:    false,
		},
		{
			name:           "successful power operation with negative base",
			operands:       []float64{-2, 3},
			precision:      2,
			expectedResult: -8,
			expectError:    false,
		},
		{
			name:           "successful power operation with negative exponent",
			operands:       []float64{2, -2},
			precision:      2,
			expectedResult: 0.25,
			expectError:    false,
		},
		{
			name:           "successful power operation with zero exponent",
			operands:       []float64{5, 0},
			precision:      2,
			expectedResult: 1,
			expectError:    false,
		},
		{
			name:           "rounding to specified precision",
			operands:       []float64{2.345, 2.789},
			precision:      2,
			expectedResult: 7.89,
			expectError:    false,
		},
		{
			name:           "rounding to zero precision",
			operands:       []float64{2.345, 2.789},
			precision:      0,
			expectedResult: 8,
			expectError:    false,
		},
		{
			name:           "error when more than two operands",
			operands:       []float64{2, 3, 4},
			precision:      2,
			expectedResult: 0,
			expectError:    true,
		},
		{
			name:           "error when less than two operands",
			operands:       []float64{2},
			precision:      2,
			expectedResult: 0,
			expectError:    true,
		},
		{
			name:           "error when no operands",
			operands:       []float64{},
			precision:      2,
			expectedResult: 0,
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			op := NewPowerOperation()
			result, err := op.Execute(tc.precision, tc.operands...)

			if tc.expectError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
