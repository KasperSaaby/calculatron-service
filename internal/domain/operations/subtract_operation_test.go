package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_SubtractOperation(t *testing.T) {
	testCases := []struct {
		name           string
		operands       []float64
		precision      int
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "successful subtraction of two numbers",
			operands:       []float64{5, 3},
			precision:      2,
			expectedResult: 2,
			expectError:    false,
		},
		{
			name:           "successful subtraction with negative result",
			operands:       []float64{3, 5},
			precision:      2,
			expectedResult: -2,
			expectError:    false,
		},
		{
			name:           "successful subtraction with decimal numbers",
			operands:       []float64{5.5, 3.2},
			precision:      2,
			expectedResult: 2.3,
			expectError:    false,
		},
		{
			name:           "successful subtraction with zero",
			operands:       []float64{5, 0},
			precision:      2,
			expectedResult: 5,
			expectError:    false,
		},
		{
			name:           "rounding to specified precision",
			operands:       []float64{5.345, 3.789},
			precision:      2,
			expectedResult: 1.56,
			expectError:    false,
		},
		{
			name:           "rounding to zero precision",
			operands:       []float64{5.345, 3.789},
			precision:      0,
			expectedResult: 2,
			expectError:    false,
		},
		{
			name:           "error when more than two operands",
			operands:       []float64{5, 3, 1},
			precision:      2,
			expectedResult: 0,
			expectError:    true,
		},
		{
			name:           "error when less than two operands",
			operands:       []float64{5},
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
			op := NewSubtractOperation()
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
