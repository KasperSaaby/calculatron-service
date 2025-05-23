package database

import (
	"database/sql"
	"embed"
	"errors"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

//go:embed migrations/*.sql
var embeddedFiles embed.FS

func MigrateSchemas(conn *sql.DB) error {
	dbInstance, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		return err
	}

	sourceDriver, err := iofs.New(embeddedFiles, "migrations")
	if err != nil {
		return err
	}

	dbName := os.Getenv("DBNAME")
	if dbName == "" {
		return errors.New("DBNAME environment variable not set")
	}

	m, err := migrate.NewWithInstance("iofs", sourceDriver, dbName, dbInstance)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
