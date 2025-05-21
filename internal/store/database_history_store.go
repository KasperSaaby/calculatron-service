package store

import (
	"context"

	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/KasperSaaby/calculatron-service/internal/store/repository"
)

type DatabaseHistoryStore struct {
	repo *repository.HistoryRepo
}

func NewDatabaseHistoryStore(repo *repository.HistoryRepo) *DatabaseHistoryStore {
	return &DatabaseHistoryStore{
		repo: repo,
	}
}

func (s *DatabaseHistoryStore) SaveCalculation(ctx context.Context, entry values.HistoryEntry) error {
	return s.repo.Create(ctx, entry)
}

func (s *DatabaseHistoryStore) GetAllCalculations(ctx context.Context, offset, limit int) ([]values.HistoryEntry, error) {
	return s.repo.FindAll(ctx, offset, limit)
}

func (s *DatabaseHistoryStore) GetCalculationByID(ctx context.Context, operationID values.OperationID) (values.HistoryEntry, error) {
	return s.repo.FindByID(ctx, operationID)
}
