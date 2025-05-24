package store

import (
	"database/sql"
	"errors"
	"os"

	"go.uber.org/fx"
)

var Setup = fx.Options(
	fx.Provide(
		func(db *sql.DB) (Factory, error) {
			storeType := os.Getenv("HISTORY_STORE_TYPE")
			if storeType == "" {
				return nil, errors.New("HISTORY_STORE_TYPE must be defined")
			}

			return GetStoreFactory(Type(storeType), db)
		},
		func(factory Factory) (HistoryStore, error) {
			return factory.CreateHistoryStore()
		},
	),
)
