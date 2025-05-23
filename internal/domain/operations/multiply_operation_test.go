package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_MultiplyOperation(t *testing.T) {
	testCases := []struct {
		name           string
		operands       []float64
		precision      int
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "successful multiplication of two numbers",
			operands:       []float64{2, 3},
			precision:      2,
			expectedResult: 6,
			expectError:    false,
		},
		{
			name:           "successful multiplication of multiple numbers",
			operands:       []float64{2, 3, 4},
			precision:      2,
			expectedResult: 24,
			expectError:    false,
		},
		{
			name:           "successful multiplication with decimal numbers",
			operands:       []float64{2.5, 3.5},
			precision:      2,
			expectedResult: 8.75,
			expectError:    false,
		},
		{
			name:           "successful multiplication with negative numbers",
			operands:       []float64{-2, 3, -1},
			precision:      2,
			expectedResult: 6,
			expectError:    false,
		},
		{
			name:           "successful multiplication with zero",
			operands:       []float64{5, 0, 3},
			precision:      2,
			expectedResult: 0,
			expectError:    false,
		},
		{
			name:           "rounding to specified precision",
			operands:       []float64{2.345, 3.789},
			precision:      2,
			expectedResult: 8.89,
			expectError:    false,
		},
		{
			name:           "rounding to zero precision",
			operands:       []float64{2.345, 3.789},
			precision:      0,
			expectedResult: 9,
			expectError:    false,
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
			op := NewMultiplyOperation()
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
