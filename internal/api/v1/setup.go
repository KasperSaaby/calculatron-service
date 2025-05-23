package v1

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/KasperSaaby/calculatron-service/generated/restapi/operations"
	"github.com/KasperSaaby/calculatron-service/internal/api/v1/handlers"
	"github.com/KasperSaaby/calculatron-service/internal/app"
	domain "github.com/KasperSaaby/calculatron-service/internal/domain/operations"
	"github.com/KasperSaaby/calculatron-service/internal/store"
)

func Setup(api *operations.CalculatronServiceAPI, db *sql.DB) error {
	storeType := os.Getenv("HISTORY_STORE_TYPE")
	if storeType == "" {
		return errors.New("HISTORY_STORE_TYPE must be defined")
	}

	storeFactory, err := store.GetStoreFactory(store.Type(storeType), db)
	if err != nil {
		return fmt.Errorf("get store factory: %w", err)
	}

	historyStore, err := storeFactory.CreateHistoryStore()
	if err != nil {
		return fmt.Errorf("create history store: %w", err)
	}

	operationFactory := domain.NewOperationFactory()
	calculatorService := app.NewCalculatorService(operationFactory)
	calculatorServiceDecorator := app.NewCalculatorServiceDecorator(calculatorService, historyStore)
	historyService := app.NewHistoryService(historyStore)

	api.GetPingHandler = handlers.GetPingHandler()
	api.PostCalculatorHandler = handlers.PostCalculateHandler(calculatorServiceDecorator)
	api.GetHistoryEntriesHandler = handlers.GetHistoryEntriesHandler(historyService)
	api.GetHistoryEntryHandler = handlers.GetHistoryEntryHandler(historyService)

	return nil
}
