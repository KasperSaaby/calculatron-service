package calculate

import (
	"bytes"
	"calculatron/internal/app/calculator"
	"calculatron/internal/domain/values"
	"calculatron/pkg/logger"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func Handler(service *calculator.Service) func(w http.ResponseWriter, r *http.Request) {
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

			var req Request
			err = json.Unmarshal(buf.Bytes(), &req)
			if err != nil {
				logger.Errf(err, "Unmarshal request body")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			result, err := service.PerformCalculation(r.Context(), values.OperationType(req.OperationType), req.Operands, req.Precision)
			if err != nil {
				var clientErr *calculator.ClientError
				if errors.As(err, &clientErr) {
					logger.Infof("Client error: %v", clientErr)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				logger.Errf(err, "Perform calculation")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			resp := Response{
				Result:      result.Result,
				Precision:   result.Precision,
				OperationID: result.OperationID,
				Timestamp:   result.Timestamp.Format(time.RFC3339),
			}

			b, err := json.Marshal(resp)
			if err != nil {
				logger.Errf(err, "Marshal response")
				w.WriteHeader(http.StatusInternalServerError)
			}

			w.Header().Set("Content-Type", "application/json")
			_, err = w.Write(b)
			if err != nil {
				logger.Errf(err, "Write response")
				w.WriteHeader(http.StatusInternalServerError)
			}

			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}
