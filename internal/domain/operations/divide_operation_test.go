package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_DivideOperation(t *testing.T) {
	testCases := []struct {
		name           string
		operands       []float64
		precision      int
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "successful division of two numbers",
			operands:       []float64{6, 2},
			precision:      2,
			expectedResult: 3,
			expectError:    false,
		},
		{
			name:           "successful division with decimal result",
			operands:       []float64{5, 2},
			precision:      2,
			expectedResult: 2.5,
			expectError:    false,
		},
		{
			name:           "successful division with negative numbers",
			operands:       []float64{-6, 2},
			precision:      2,
			expectedResult: -3,
			expectError:    false,
		},
		{
			name:           "successful division with decimal numbers",
			operands:       []float64{5.5, 2.2},
			precision:      2,
			expectedResult: 2.5,
			expectError:    false,
		},
		{
			name:           "rounding to specified precision",
			operands:       []float64{5.345, 2.789},
			precision:      2,
			expectedResult: 1.92,
			expectError:    false,
		},
		{
			name:           "rounding to zero precision",
			operands:       []float64{5.345, 2.789},
			precision:      0,
			expectedResult: 2,
			expectError:    false,
		},
		{
			name:           "error when dividing by zero",
			operands:       []float64{5, 0},
			precision:      2,
			expectedResult: 0,
			expectError:    true,
		},
		{
			name:           "error when more than two operands",
			operands:       []float64{6, 2, 1},
			precision:      2,
			expectedResult: 0,
			expectError:    true,
		},
		{
			name:           "error when less than two operands",
			operands:       []float64{6},
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
			op := NewDivideOperation()
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
