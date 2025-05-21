package v1

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/KasperSaaby/calculatron-service/internal/api/v1/handlers/calculate"
	"github.com/KasperSaaby/calculatron-service/internal/api/v1/handlers/ping"
	"github.com/KasperSaaby/calculatron-service/internal/app"
	"github.com/KasperSaaby/calculatron-service/internal/store"
)

func Setup(mux *http.ServeMux, db *sql.DB) error {
	storeType := os.Getenv("HISTORY_STORE_TYPE")
	if storeType == "" {
		return errors.New("HISTORY_STORE_TYPE must be set")
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

	mux.HandleFunc("/v1/ping", ping.Handler())
	mux.HandleFunc("/v1/calculate", calculate.Handler(calculatorService))

	return nil
}
