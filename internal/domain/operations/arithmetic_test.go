package operations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Add(t *testing.T) {
	testCases := []struct {
		name           string
		operands       []float64
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "successful addition of two numbers",
			operands:       []float64{2, 3},
			expectedResult: 5,
			expectError:    false,
		},
		{
			name:           "successful addition of multiple numbers",
			operands:       []float64{2, 3, 4, 5},
			expectedResult: 14,
			expectError:    false,
		},
		{
			name:           "error when less than two operands",
			operands:       []float64{1},
			expectedResult: 0,
			expectError:    true,
		},
		{
			name:           "error when no operands",
			operands:       []float64{},
			expectedResult: 0,
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Add(tc.operands...)

			if tc.expectError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func Test_Subtract(t *testing.T) {
	testCases := []struct {
		name           string
		operands       []float64
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "successful subtraction",
			operands:       []float64{5, 3},
			expectedResult: 2,
			expectError:    false,
		},
		{
			name:           "error when more than two operands",
			operands:       []float64{5, 3, 1},
			expectedResult: 0,
			expectError:    true,
		},
		{
			name:           "error when less than two operands",
			operands:       []float64{5},
			expectedResult: 0,
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Subtract(tc.operands...)

			if tc.expectError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func Test_Multiply(t *testing.T) {
	testCases := []struct {
		name           string
		operands       []float64
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "successful multiplication of two numbers",
			operands:       []float64{2, 3},
			expectedResult: 6,
			expectError:    false,
		},
		{
			name:           "successful multiplication of multiple numbers",
			operands:       []float64{2, 3, 4},
			expectedResult: 24,
			expectError:    false,
		},
		{
			name:           "error when less than two operands",
			operands:       []float64{2},
			expectedResult: 0,
			expectError:    true,
		},
		{
			name:           "error when no operands",
			operands:       []float64{},
			expectedResult: 0,
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Multiply(tc.operands...)

			if tc.expectError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func Test_Divide(t *testing.T) {
	testCases := []struct {
		name           string
		operands       []float64
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "successful division",
			operands:       []float64{6, 2},
			expectedResult: 3,
			expectError:    false,
		},
		{
			name:           "error when dividing by zero",
			operands:       []float64{6, 0},
			expectedResult: 0,
			expectError:    true,
		},
		{
			name:           "error when more than two operands",
			operands:       []float64{6, 2, 1},
			expectedResult: 0,
			expectError:    true,
		},
		{
			name:           "error when less than two operands",
			operands:       []float64{6},
			expectedResult: 0,
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Divide(tc.operands...)

			if tc.expectError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func Test_Power(t *testing.T) {
	testCases := []struct {
		name           string
		operands       []float64
		expectedResult float64
		expectError    bool
	}{
		{
			name:           "successful power operation",
			operands:       []float64{2, 3},
			expectedResult: 8,
			expectError:    false,
		},
		{
			name:           "successful power operation with decimal exponent",
			operands:       []float64{4, 0.5},
			expectedResult: 2,
			expectError:    false,
		},
		{
			name:           "error when more than two operands",
			operands:       []float64{2, 3, 4},
			expectedResult: 0,
			expectError:    true,
		},
		{
			name:           "error when less than two operands",
			operands:       []float64{2},
			expectedResult: 0,
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Power(tc.operands...)

			if tc.expectError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
