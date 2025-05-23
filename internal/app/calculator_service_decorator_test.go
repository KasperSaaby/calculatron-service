package app

import (
	"context"
	"testing"

	"github.com/KasperSaaby/calculatron-service/internal/app/models"
	"github.com/KasperSaaby/calculatron-service/internal/domain/operations"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/KasperSaaby/calculatron-service/internal/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_CalculatorServiceDecorator_PerformCalculation(t *testing.T) {
	operationFactory := operations.NewOperationFactory()
	storeFactory, err := store.GetStoreFactory(store.InMemory_Type, nil)
	require.NoError(t, err)

	t.Run("successful addition and save to history", func(t *testing.T) {
		ctx := context.Background()
		historyStore, err := storeFactory.CreateHistoryStore()
		require.NoError(t, err)
		calculatorService := NewCalculatorService(operationFactory)
		sut := NewCalculatorServiceDecorator(calculatorService, historyStore)

		operationType := values.OperationType_Add
		operands := []float64{1, 2}
		precision := 2
		expectedResult := 3.00

		input, err := models.NewCalculationInput(operationType.String(), operands, precision)
		require.NoError(t, err)

		result, err := sut.PerformCalculation(ctx, input)

		// Assert that calculation was performed successfully
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result.Result)
		assert.Equal(t, precision, result.Precision)
		assert.NotEmpty(t, result.OperationID)
		assert.NotZero(t, result.Timestamp)

		// Assert that calculation was stored in history
		history, err := historyStore.GetAllCalculations(ctx, 0, 1)
		require.NoError(t, err)
		require.Len(t, history, 1)
		assert.Equal(t, result.OperationID, history[0].OperationID)
		assert.Equal(t, operationType, history[0].OperationType)
		assert.Equal(t, operands, history[0].Operands)
		assert.Equal(t, precision, history[0].Precision)
		assert.Equal(t, expectedResult, history[0].Result)
	})
}
