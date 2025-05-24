package v1

import (
	"github.com/KasperSaaby/calculatron-service/generated/restapi/operations"
	"github.com/KasperSaaby/calculatron-service/internal/api/v1/handlers"
	"github.com/KasperSaaby/calculatron-service/internal/app"
	"go.uber.org/fx"
)

var Setup = fx.Options(
	fx.Invoke(
		func(
			api *operations.CalculatronServiceAPI,
			calculatorServiceDecorator *app.CalculatorServiceDecorator,
			historyService *app.HistoryService,
		) {
			api.GetPingHandler = handlers.GetPingHandler()
			api.PostCalculatorHandler = handlers.PostCalculateHandler(calculatorServiceDecorator)
			api.GetHistoryEntriesHandler = handlers.GetHistoryEntriesHandler(historyService)
			api.GetHistoryEntryHandler = handlers.GetHistoryEntryHandler(historyService)
		},
	),
)
