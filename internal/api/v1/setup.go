package v1

import (
	"calculatron/internal/api/v1/handlers/calculate"
	"calculatron/internal/api/v1/handlers/ping"
	"calculatron/internal/app/calculator"
	"net/http"
)

func Setup(mux *http.ServeMux) error {
	service := calculator.NewService()
	mux.HandleFunc("/v1/ping", ping.Handler())
	mux.HandleFunc("/v1/calculate", calculate.Handler(service))
	return nil
}
