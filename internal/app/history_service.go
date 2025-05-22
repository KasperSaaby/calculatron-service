package app

import (
	"context"
	"fmt"

	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/KasperSaaby/calculatron-service/internal/store"
)

type HistoryService struct {
	historyStore store.HistoryStore
}

func NewHistoryService(historyStore store.HistoryStore) *HistoryService {
	return &HistoryService{
		historyStore: historyStore,
	}
}

func (s *HistoryService) GetHistory(ctx context.Context, offset, limit int) ([]values.HistoryEntry, error) {
	historyEntries, err := s.historyStore.GetAllCalculations(ctx, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("get history entries: %w", err)
	}

	return historyEntries, nil
}

func (s *HistoryService) GetHistoryByID(ctx context.Context, operationID values.OperationID) (values.HistoryEntry, error) {
	historyEntry, err := s.historyStore.GetCalculationByID(ctx, operationID)
	if err != nil {
		return values.HistoryEntry{}, fmt.Errorf("get history entry: %w", err)
	}

	return historyEntry, nil
}
