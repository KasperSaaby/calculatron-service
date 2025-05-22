package v1

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/KasperSaaby/calculatron-service/generated/restapi/operations"
	"github.com/KasperSaaby/calculatron-service/internal/api/v1/handlers/calculate"
	"github.com/KasperSaaby/calculatron-service/internal/api/v1/handlers/history"
	"github.com/KasperSaaby/calculatron-service/internal/api/v1/handlers/ping"
	"github.com/KasperSaaby/calculatron-service/internal/app"
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

	calculatorService := app.NewCalculatorService(historyStore)
	historyService := app.NewHistoryService(historyStore)

	api.GetPingHandler = ping.GetPingHandler()
	api.PostCalculatorHandler = calculate.PostCalculateHandler(calculatorService)
	api.GetHistoryEntriesHandler = history.GetHistoryEntriesHandler(historyService)
	api.GetHistoryEntryHandler = history.GetHistoryEntryHandler(historyService)

	return nil
}
