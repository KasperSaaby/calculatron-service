package v1

import (
	"database/sql"
	"net/http"

	"github.com/KasperSaaby/calculatron-service/internal/api/v1/handlers/calculate"
	"github.com/KasperSaaby/calculatron-service/internal/api/v1/handlers/ping"
	"github.com/KasperSaaby/calculatron-service/internal/app"
	"github.com/KasperSaaby/calculatron-service/internal/store/database/repos"
)

func Setup(mux *http.ServeMux, db *sql.DB) error {
	var (
		calculationHistoryRepo = repos.NewCalculationHistoryRepo(db)
		service                = app.NewCalculatorService(calculationHistoryRepo)
	)

	mux.HandleFunc("/v1/ping", ping.Handler())
	mux.HandleFunc("/v1/calculate", calculate.Handler(service))

	return nil
}
