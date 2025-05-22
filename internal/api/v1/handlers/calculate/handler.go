package calculate

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/KasperSaaby/calculatron-service/internal/app"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/KasperSaaby/calculatron-service/internal/platform/logger"
)

func Handler(service *app.CalculatorService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			buf := new(bytes.Buffer)
			_, err := buf.ReadFrom(r.Body)
			if err != nil {
				logger.Errf(err, "Read request body")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var req PostCalculateRequest
			err = json.Unmarshal(buf.Bytes(), &req)
			if err != nil {
				logger.Errf(err, "Unmarshal request body")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			result, err := service.PerformCalculation(
				r.Context(),
				values.OperationType(req.OperationType),
				req.Operands,
				req.Precision,
			)
			if err != nil {
				var appError *app.AppError
				if errors.As(err, &appError) {
					logger.Infof("App error: %v", appError)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				logger.Errf(err, "Perform calculation")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			resp := PostCalculateResponse{
				Result:      result.Result,
				Precision:   result.Precision,
				OperationID: result.OperationID,
				Timestamp:   result.Timestamp.Format(time.RFC3339),
			}

			b, err := json.Marshal(resp)
			if err != nil {
				logger.Errf(err, "Marshal response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write(b)
			if err != nil {
				logger.Errf(err, "Write response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}
