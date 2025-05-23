package handlers

import (
	"errors"
	"time"

	"github.com/KasperSaaby/calculatron-service/generated/models"
	"github.com/KasperSaaby/calculatron-service/generated/restapi/operations"
	"github.com/KasperSaaby/calculatron-service/internal/app"
	appmodels "github.com/KasperSaaby/calculatron-service/internal/app/models"
	"github.com/KasperSaaby/calculatron-service/internal/platform/logger"
	"github.com/go-openapi/runtime/middleware"
)

func PostCalculateHandler(calculatorService app.Calculator) operations.PostCalculatorHandlerFunc {
	return func(params operations.PostCalculatorParams) middleware.Responder {
		input := appmodels.NewCalculationInput(
			params.Body.OperationType,
			params.Body.Operands,
			int(params.Body.Precision),
		)

		result, err := calculatorService.PerformCalculation(params.HTTPRequest.Context(), input)
		if err != nil {
			var appError *app.AppError
			if errors.As(err, &appError) {
				logger.Infof("App error: %v", appError)
				return operations.NewPostCalculatorBadRequest()
			}

			logger.Errf(err, "Perform calculation")
			return operations.NewPostCalculatorInternalServerError()
		}

		return operations.NewPostCalculatorOK().WithPayload(&models.PostCalculateResponse{
			OperationID: result.OperationID.String(),
			Precision:   int32(result.Precision),
			Result:      result.Result,
			Timestamp:   result.Timestamp.Format(time.RFC3339),
		})
	}
}
