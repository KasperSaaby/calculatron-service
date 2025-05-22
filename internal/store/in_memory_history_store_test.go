package store

import (
	"context"
	"testing"
	"time"

	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_InMemoryHistoryStore_SaveCalculation(t *testing.T) {
	duplicateOperationID := values.NewOperationID()

	testCases := []struct {
		name          string
		newEntry      values.HistoryEntry
		existingEntry *values.HistoryEntry
		expectError   bool
	}{
		{
			name: "successfully save new entry",
			newEntry: values.HistoryEntry{
				OperationID:   values.NewOperationID(),
				OperationType: values.OperationType_Add,
				Operands:      []float64{1, 2},
				Result:        3,
				Precision:     2,
				Timestamp:     time.Now(),
			},
			expectError: false,
		},
		{
			name: "error when saving duplicate entry",
			newEntry: values.HistoryEntry{
				OperationID:   duplicateOperationID,
				OperationType: values.OperationType_Add,
				Operands:      []float64{1, 2},
				Result:        3,
				Precision:     2,
				Timestamp:     time.Now(),
			},
			existingEntry: &values.HistoryEntry{
				OperationID:   duplicateOperationID,
				OperationType: values.OperationType_Add,
				Operands:      []float64{1, 2},
				Result:        3,
				Precision:     2,
				Timestamp:     time.Now(),
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := NewInMemoryHistoryStore()
			ctx := context.Background()

			if tc.existingEntry != nil {
				err := sut.SaveCalculation(ctx, *tc.existingEntry)
				require.NoError(t, err)
			}

			err := sut.SaveCalculation(ctx, tc.newEntry)

			if tc.expectError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			entries, err := sut.GetAllCalculations(ctx, 0, 1)
			assert.NoError(t, err)
			assert.Len(t, entries, 1)
			assert.Equal(t, tc.newEntry, entries[0])
		})
	}
}

func Test_InMemoryHistoryStore_GetAllCalculations(t *testing.T) {
	testCases := []struct {
		name            string
		existingEntries []values.HistoryEntry
		offset          int
		limit           int
		expectedCount   int
		expectedError   bool
	}{
		{
			name:            "empty store returns empty slice",
			existingEntries: []values.HistoryEntry{},
			offset:          0,
			limit:           10,
			expectedCount:   0,
			expectedError:   false,
		},
		{
			name: "returns all entries within limit",
			existingEntries: []values.HistoryEntry{
				{
					OperationID:   values.NewOperationID(),
					OperationType: values.OperationType_Add,
					Operands:      []float64{1, 2},
					Result:        3,
					Precision:     2,
					Timestamp:     time.Now(),
				},
				{
					OperationID:   values.NewOperationID(),
					OperationType: values.OperationType_Multiply,
					Operands:      []float64{2, 3},
					Result:        6,
					Precision:     2,
					Timestamp:     time.Now(),
				},
			},
			offset:        0,
			limit:         10,
			expectedCount: 2,
			expectedError: false,
		},
		{
			name: "respects limit",
			existingEntries: []values.HistoryEntry{
				{
					OperationID:   values.NewOperationID(),
					OperationType: values.OperationType_Add,
					Operands:      []float64{1, 2},
					Result:        3,
					Precision:     2,
					Timestamp:     time.Now(),
				},
				{
					OperationID:   values.NewOperationID(),
					OperationType: values.OperationType_Multiply,
					Operands:      []float64{2, 3},
					Result:        6,
					Precision:     2,
					Timestamp:     time.Now(),
				},
			},
			offset:        0,
			limit:         1,
			expectedCount: 1,
			expectedError: false,
		},
		{
			name: "respects offset",
			existingEntries: []values.HistoryEntry{
				{
					OperationID:   values.NewOperationID(),
					OperationType: values.OperationType_Add,
					Operands:      []float64{1, 2},
					Result:        3,
					Precision:     2,
					Timestamp:     time.Now(),
				},
				{
					OperationID:   values.NewOperationID(),
					OperationType: values.OperationType_Multiply,
					Operands:      []float64{2, 3},
					Result:        6,
					Precision:     2,
					Timestamp:     time.Now(),
				},
			},
			offset:        1,
			limit:         1,
			expectedCount: 1,
			expectedError: false,
		},
		{
			name: "negative offset is treated as 0",
			existingEntries: []values.HistoryEntry{
				{
					OperationID:   values.NewOperationID(),
					OperationType: values.OperationType_Add,
					Operands:      []float64{1, 2},
					Result:        3,
					Precision:     2,
					Timestamp:     time.Now(),
				},
			},
			offset:        -1,
			limit:         10,
			expectedCount: 1,
			expectedError: false,
		},
		{
			name: "zero or negative limit returns empty slice",
			existingEntries: []values.HistoryEntry{
				{
					OperationID:   values.NewOperationID(),
					OperationType: values.OperationType_Add,
					Operands:      []float64{1, 2},
					Result:        3,
					Precision:     2,
					Timestamp:     time.Now(),
				},
			},
			offset:        0,
			limit:         0,
			expectedCount: 0,
			expectedError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := NewInMemoryHistoryStore()
			ctx := context.Background()

			for _, entry := range tc.existingEntries {
				err := sut.SaveCalculation(ctx, entry)
				require.NoError(t, err)
			}

			entries, err := sut.GetAllCalculations(ctx, tc.offset, tc.limit)

			if tc.expectedError {
				assert.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Len(t, entries, tc.expectedCount)
		})
	}
}

func Test_InMemoryHistoryStore_GetCalculationByID(t *testing.T) {
	operationsID := values.NewOperationID()

	testCases := []struct {
		name          string
		existingEntry values.HistoryEntry
		operationID   values.OperationID
		expectError   bool
		expectedEntry values.HistoryEntry
	}{
		{
			name: "successfully retrieve existing entry",
			existingEntry: values.HistoryEntry{
				OperationID:   operationsID,
				OperationType: values.OperationType_Add,
				Operands:      []float64{1, 2},
				Result:        3,
				Precision:     2,
				Timestamp:     time.Now(),
			},
			operationID:   operationsID,
			expectError:   false,
			expectedEntry: values.HistoryEntry{},
		},
		{
			name: "error when entry not found",
			existingEntry: values.HistoryEntry{
				OperationID:   values.NewOperationID(),
				OperationType: values.OperationType_Add,
				Operands:      []float64{1, 2},
				Result:        3,
				Precision:     2,
				Timestamp:     time.Now(),
			},
			operationID:   values.NewOperationID(),
			expectError:   true,
			expectedEntry: values.HistoryEntry{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sut := NewInMemoryHistoryStore()
			ctx := context.Background()

			err := sut.SaveCalculation(ctx, tc.existingEntry)
			require.NoError(t, err)

			entry, err := sut.GetCalculationByID(ctx, tc.operationID)

			if tc.expectError {
				assert.Error(t, err)
				assert.Equal(t, values.ErrHistoryEntryNotFound, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.existingEntry, entry)
		})
	}
}
