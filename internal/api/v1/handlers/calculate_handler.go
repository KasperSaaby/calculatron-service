package handlers

import (
	"errors"
	"time"

	"github.com/KasperSaaby/calculatron-service/generated/models"
	"github.com/KasperSaaby/calculatron-service/generated/restapi/operations"
	"github.com/KasperSaaby/calculatron-service/internal/app"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/KasperSaaby/calculatron-service/internal/platform/logger"
	"github.com/go-openapi/runtime/middleware"
)

func PostCalculateHandler(service *app.CalculatorService) operations.PostCalculatorHandlerFunc {
	return func(params operations.PostCalculatorParams) middleware.Responder {
		result, err := service.PerformCalculation(
			params.HTTPRequest.Context(),
			values.OperationType(params.Body.OperationType),
			params.Body.Operands,
			int(params.Body.Precision),
		)
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
