package store

import (
	"context"
	"fmt"
	"sync"

	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
)

type InMemoryHistoryStore struct {
	mux     sync.RWMutex
	entries map[values.OperationID]values.HistoryEntry
	order   []values.OperationID
}

func NewInMemoryHistoryStore() *InMemoryHistoryStore {
	return &InMemoryHistoryStore{
		entries: make(map[values.OperationID]values.HistoryEntry),
		order:   make([]values.OperationID, 0),
	}
}

func (s *InMemoryHistoryStore) SaveCalculation(_ context.Context, entry values.HistoryEntry) error {
	s.mux.Lock()
	defer s.mux.Unlock()

	_, exist := s.entries[entry.OperationID]
	if exist {
		return fmt.Errorf("entry with operation id %v already exists", entry.OperationID)
	}

	s.entries[entry.OperationID] = entry
	s.order = append(s.order, entry.OperationID)

	return nil
}

func (s *InMemoryHistoryStore) GetAllCalculations(_ context.Context, offset, limit int) ([]values.HistoryEntry, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	if offset < 0 {
		offset = 0
	}

	if limit <= 0 {
		return []values.HistoryEntry{}, nil
	}

	totalEntries := len(s.order)
	if offset >= totalEntries {
		return []values.HistoryEntry{}, nil
	}

	endIndex := offset + limit
	if endIndex > totalEntries {
		endIndex = totalEntries
	}

	var entries []values.HistoryEntry
	for _, id := range s.order[offset:endIndex] {
		entries = append(entries, s.entries[id])
	}

	return entries, nil
}

func (s *InMemoryHistoryStore) GetCalculationByID(_ context.Context, operationID values.OperationID) (values.HistoryEntry, error) {
	s.mux.RLock()
	defer s.mux.RUnlock()

	entry, exist := s.entries[operationID]
	if !exist {
		return values.HistoryEntry{}, values.ErrHistoryEntryNotFound
	}

	return entry, nil
}
