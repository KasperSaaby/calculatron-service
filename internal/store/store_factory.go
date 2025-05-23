package store

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/KasperSaaby/calculatron-service/internal/domain/values"
	"github.com/KasperSaaby/calculatron-service/internal/store/in_memory"
	"github.com/KasperSaaby/calculatron-service/internal/store/repository"
)

type Factory interface {
	CreateHistoryStore() (HistoryStore, error)
}

//go:generate moq -pkg mocks -out mocks/history_store_mock.go . HistoryStore
type HistoryStore interface {
	SaveCalculation(ctx context.Context, entry values.HistoryEntry) error
	GetAllCalculations(ctx context.Context, offset, limit int) ([]values.HistoryEntry, error)
	GetCalculationByID(ctx context.Context, operationID values.OperationID) (values.HistoryEntry, error)
}

type databaseStoreFactory struct {
	db *sql.DB
}

func NewDatabaseStoreFactory(db *sql.DB) (Factory, error) {
	if db == nil {
		return nil, fmt.Errorf("database connection is required for database store factory")
	}
	return &databaseStoreFactory{db: db}, nil
}

func (f *databaseStoreFactory) CreateHistoryStore() (HistoryStore, error) {
	return NewDatabaseHistoryStore(repository.NewHistoryRepo(f.db)), nil
}

type inMemoryStoreFactory struct {
}

func NewInMemoryStoreFactory() (Factory, error) {
	return &inMemoryStoreFactory{}, nil
}

func (f *inMemoryStoreFactory) CreateHistoryStore() (HistoryStore, error) {
	return NewInMemoryHistoryStore(in_memory.NewHistoryStore()), nil
}

type Type string

const (
	InMemory_Type Type = "in_memory"
	Database_Type Type = "database"
)

func GetStoreFactory(storeType Type, db *sql.DB) (Factory, error) {
	switch storeType {
	case InMemory_Type:
		return NewInMemoryStoreFactory()
	case Database_Type:
		return NewDatabaseStoreFactory(db)
	default:
		return nil, fmt.Errorf("unknown type: %s", storeType)
	}
}
