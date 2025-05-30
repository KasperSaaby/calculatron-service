package handlers

import (
	"errors"
	"time"

	"github.com/KasperSaaby/calculatron-service/generated/models"
	"github.com/KasperSaaby/calculatron-service/generated/restapi/operations"
	"github.com/KasperSaaby/calculatron-service/internal/app"
	"github.com/KasperSaaby/calculatron-service/internal/platform/logger"
	"github.com/go-openapi/runtime/middleware"
)

const (
	QueryParam_DefaultLimit = 20
	QueryParam_MaxLimit     = 100
)

func GetHistoryEntriesHandler(historyService *app.HistoryService) operations.GetHistoryEntriesHandlerFunc {
	return func(params operations.GetHistoryEntriesParams) middleware.Responder {
		offset := int(params.Offset)
		if offset < 0 {
			offset = 0
		}

		limit := int(params.Limit)
		if limit < 0 {
			limit = QueryParam_DefaultLimit
		}

		if limit > QueryParam_MaxLimit {
			limit = QueryParam_MaxLimit
		}

		ctx := params.HTTPRequest.Context()
		historyEntries, err := historyService.GetHistory(ctx, offset, limit)
		if err != nil {
			logger.Errf(ctx, err, "Get history entries")
			return operations.NewGetHistoryEntriesInternalServerError()
		}

		resp := &models.GetHistoryEntriesResponse{}
		for _, entry := range historyEntries {
			resp.Entries = append(resp.Entries, &models.Entry{
				Operands:      entry.Operands,
				OperationID:   entry.OperationID.String(),
				OperationType: entry.OperationType.String(),
				Precision:     int32(entry.Precision),
				Result:        entry.Result,
				Timestamp:     entry.Timestamp.Format(time.RFC3339),
			})
		}

		return operations.NewGetHistoryEntriesOK().WithPayload(resp)
	}
}

func GetHistoryEntryHandler(historyService *app.HistoryService) operations.GetHistoryEntryHandlerFunc {
	return func(params operations.GetHistoryEntryParams) middleware.Responder {
		ctx := params.HTTPRequest.Context()
		entry, err := historyService.GetHistoryByID(ctx, params.OperationID)
		if err != nil {
			if errors.Is(err, app.ErrNotFound) {
				resp := &models.ErrorModel{
					Message:    "No history entry found",
					ReasonCode: "not_found",
				}

				return operations.NewGetHistoryEntryNotFound().WithPayload(resp)
			}

			logger.Errf(ctx, err, "Get history entry")
			return operations.NewGetHistoryEntryInternalServerError()
		}

		resp := &models.GetHistoryEntryResponse{
			Entry: &models.Entry{
				Operands:      entry.Operands,
				OperationID:   entry.OperationID.String(),
				OperationType: entry.OperationType.String(),
				Precision:     int32(entry.Precision),
				Result:        entry.Result,
				Timestamp:     entry.Timestamp.Format(time.RFC3339),
			},
		}

		return operations.NewGetHistoryEntryOK().WithPayload(resp)
	}
}
