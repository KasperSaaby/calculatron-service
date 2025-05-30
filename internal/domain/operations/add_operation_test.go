package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_AddOperation(t *testing.T) {
	testCases := []struct {
		name           string
		operands       []float64
		precision      int
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "successful addition of two numbers",
			operands:       []float64{2, 3},
			precision:      2,
			expectedResult: 5,
			expectError:    false,
		},
		{
			name:           "successful addition of multiple numbers",
			operands:       []float64{2, 3, 4, 5},
			precision:      2,
			expectedResult: 14,
			expectError:    false,
		},
		{
			name:           "error when less than two operands",
			operands:       []float64{1},
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
		{
			name:           "successful addition with decimal numbers",
			operands:       []float64{2.5, 3.7},
			precision:      2,
			expectedResult: 6.2,
			expectError:    false,
		},
		{
			name:           "successful addition with negative numbers",
			operands:       []float64{-2, 3, -1},
			precision:      2,
			expectedResult: 0,
			expectError:    false,
		},
		{
			name:           "successful addition with zero",
			operands:       []float64{0, 5, 0},
			precision:      2,
			expectedResult: 5,
			expectError:    false,
		},
		{
			name:           "rounding to specified precision",
			operands:       []float64{2.345, 3.789},
			precision:      2,
			expectedResult: 6.13,
			expectError:    false,
		},
		{
			name:           "rounding to zero precision",
			operands:       []float64{2.345, 3.789},
			precision:      0,
			expectedResult: 6,
			expectError:    false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			op := NewAddOperation()
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
