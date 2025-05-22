package app

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/KasperSaaby/calculatron-service/internal/store/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_HistoryStore_GetHistory(t *testing.T) {
	testCases := []struct {
		name          string
		offset        int
		limit         int
		mockEntries   []values.HistoryEntry
		mockError     error
		expectedError bool
		expectedCount int
	}{
		{
			name:          "successful retrieval with empty history",
			offset:        0,
			limit:         10,
			mockEntries:   []values.HistoryEntry{},
			mockError:     nil,
			expectedError: false,
			expectedCount: 0,
		},
		{
			name:   "successful retrieval with entries",
			offset: 0,
			limit:  10,
			mockEntries: []values.HistoryEntry{
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
			mockError:     nil,
			expectedError: false,
			expectedCount: 2,
		},
		{
			name:          "store error",
			offset:        0,
			limit:         10,
			mockEntries:   nil,
			mockError:     errors.New("database error"),
			expectedError: true,
			expectedCount: 0,
		},
		{
			name:   "pagination with offset",
			offset: 1,
			limit:  1,
			mockEntries: []values.HistoryEntry{
				{
					OperationID:   values.NewOperationID(),
					OperationType: values.OperationType_Add,
					Operands:      []float64{1, 2},
					Result:        3,
					Precision:     2,
					Timestamp:     time.Now(),
				},
			},
			mockError:     nil,
			expectedError: false,
			expectedCount: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockStore := &mocks.HistoryStoreMock{}
			sut := NewHistoryService(mockStore)

			mockStore.GetAllCalculationsFunc = func(ctx context.Context, offset, limit int) ([]values.HistoryEntry, error) {
				assert.Equal(t, tc.offset, offset)
				assert.Equal(t, tc.limit, limit)
				return tc.mockEntries, tc.mockError
			}

			history, err := sut.GetHistory(ctx, tc.offset, tc.limit)

			if tc.expectedError {
				assert.Error(t, err)
				assert.Nil(t, history)
				return
			}

			require.NoError(t, err)
			assert.Len(t, history, tc.expectedCount)
			assert.Equal(t, tc.mockEntries, history)
		})
	}
}

func Test_HistoryStore_GetHistoryByID(t *testing.T) {
	testCases := []struct {
		name          string
		operationID   values.OperationID
		mockEntry     values.HistoryEntry
		mockError     error
		expectedError bool
	}{
		{
			name:        "successful retrieval",
			operationID: values.NewOperationID(),
			mockEntry: values.HistoryEntry{
				OperationID:   values.NewOperationID(),
				OperationType: values.OperationType_Add,
				Operands:      []float64{1, 2},
				Result:        3,
				Precision:     2,
				Timestamp:     time.Now(),
			},
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "store error",
			operationID:   values.NewOperationID(),
			mockEntry:     values.HistoryEntry{},
			mockError:     errors.New("database error"),
			expectedError: true,
		},
		{
			name:          "not found error",
			operationID:   values.NewOperationID(),
			mockEntry:     values.HistoryEntry{},
			mockError:     values.ErrHistoryEntryNotFound,
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			mockStore := &mocks.HistoryStoreMock{}
			sut := NewHistoryService(mockStore)

			mockStore.GetCalculationByIDFunc = func(ctx context.Context, operationID values.OperationID) (values.HistoryEntry, error) {
				assert.Equal(t, tc.operationID, operationID)
				return tc.mockEntry, tc.mockError
			}

			entry, err := sut.GetHistoryByID(ctx, tc.operationID)

			if tc.expectedError {
				assert.Error(t, err)
				assert.Empty(t, entry)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.mockEntry, entry)
		})
	}
}
