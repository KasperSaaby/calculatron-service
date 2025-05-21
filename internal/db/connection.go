package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func New() (*sql.DB, error) {
	dbName := os.Getenv("DBNAME")
	dbHost := os.Getenv("DBHOST")
	dbPassword := os.Getenv("DBPASSWORD")
	dbUser := os.Getenv("DBUSER")

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s", dbUser, dbPassword, dbName, dbHost)

	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("open db connection to database: %w", err)
	}

	return db, nil
}
