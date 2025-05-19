package v1

import (
	"calculatron/internal/api/v1/handlers/calculate"
	"calculatron/internal/api/v1/handlers/ping"
	"calculatron/internal/app/calculator"
	"calculatron/internal/db/repos"
	"database/sql"
	"net/http"
)

func Setup(mux *http.ServeMux, db *sql.DB) error {
	var (
		calculationHistoryRepo = repos.NewCalculationHistoryRepo(db)
		service                = calculator.NewService(calculationHistoryRepo)
	)

	mux.HandleFunc("/v1/ping", ping.Handler())
	mux.HandleFunc("/v1/calculate", calculate.Handler(service))

	return nil
}
