package history

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/KasperSaaby/calculatron-service/internal/app"
	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/KasperSaaby/calculatron-service/internal/platform/logger"
)

const (
	QueryParam_DefaultLimit = 20
	QueryParam_MaxLimit     = 100
)

func Handler(historyService *app.HistoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			offsetStr := r.URL.Query().Get("offset")
			offset, err := strconv.Atoi(offsetStr)
			if err != nil {
				logger.Infof("Invalid offset parameter %q: %v", offsetStr, err)
				offset = 0
			}

			if offset < 0 {
				offset = 0
			}

			limitStr := r.URL.Query().Get("limit")
			limit, err := strconv.Atoi(limitStr)
			if err != nil {
				logger.Infof("Invalid limit parameter %q: %v", limitStr, err)
				limit = QueryParam_DefaultLimit
			}

			if limit < 0 {
				limit = 20
			}

			if limit > QueryParam_MaxLimit {
				limit = QueryParam_MaxLimit
			}

			historyEntries, err := historyService.GetHistory(r.Context(), offset, limit)
			if err != nil {
				logger.Errf(err, "Get history entries")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var resp GetHistoryEntriesResponse
			for _, entry := range historyEntries {
				resp.Entries = append(resp.Entries, Entry{
					OperationID:   entry.OperationID.String(),
					OperationType: entry.OperationType.String(),
					Operands:      entry.Operands,
					Result:        entry.Result,
					Precision:     entry.Precision,
					Timestamp:     entry.Timestamp.Format(time.RFC3339),
				})
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
		}
	}
}

func HandlerWithPathID(historyService *app.HistoryService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			operationId := r.PathValue("operationId")
			if operationId == "" {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			historyEntry, err := historyService.GetHistoryByID(r.Context(), values.OperationID(operationId))
			if err != nil {
				logger.Errf(err, "Get history entry")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			resp := GetHistoryEntryResponse{
				Entry: Entry{
					OperationID:   historyEntry.OperationID.String(),
					OperationType: historyEntry.OperationType.String(),
					Operands:      historyEntry.Operands,
					Result:        historyEntry.Result,
					Precision:     historyEntry.Precision,
					Timestamp:     historyEntry.Timestamp.Format(time.RFC3339),
				},
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
		}
	}
}
