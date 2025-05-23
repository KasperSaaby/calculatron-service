package store

import (
	"context"

	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/KasperSaaby/calculatron-service/internal/store/in_memory"
)

type InMemoryHistoryStore struct {
	inMemoryStore *in_memory.HistoryStore
}

func NewInMemoryHistoryStore(inMemoryStore *in_memory.HistoryStore) *InMemoryHistoryStore {
	return &InMemoryHistoryStore{
		inMemoryStore: inMemoryStore,
	}
}

func (s *InMemoryHistoryStore) SaveCalculation(_ context.Context, entry values.HistoryEntry) error {
	return s.inMemoryStore.Create(entry)
}

func (s *InMemoryHistoryStore) GetAllCalculations(_ context.Context, offset, limit int) ([]values.HistoryEntry, error) {
	return s.inMemoryStore.GetAll(offset, limit)
}

func (s *InMemoryHistoryStore) GetCalculationByID(_ context.Context, operationID values.OperationID) (values.HistoryEntry, error) {
	return s.inMemoryStore.GetByID(operationID)
}
