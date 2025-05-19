package db

import (
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/stdlib"
)

func New() (*sql.DB, error) {
	userInfo := url.UserPassword("postgres", "password")

	dns := fmt.Sprintf(
		"postgresql://%s@%s:%d/%s?%s%s",
		userInfo.String(),
		"localhost",
		5432,
		"calculatron",
		sslModeValue(false),
		searchPath(""),
	)

	connector, err := GetConnector(stdlib.GetDefaultDriver(), dns)
	if err != nil {
		return nil, err
	}

	return OpenDB(connector)
}

func GetConnector(dri driver.Driver, dsn string) (driver.Connector, error) {
	driverContext, ok := dri.(driver.DriverContext)
	if ok {
		return driverContext.OpenConnector(dsn)
	}

	return nil, errors.New("driver does not implement driver.DriverContext")
}

func OpenDB(connector driver.Connector) (*sql.DB, error) {
	db := sql.OpenDB(connector)
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(2)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(0)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func sslModeValue(requireTLS bool) string {
	if requireTLS {
		return "sslmode=require"
	}

	return "sslmode=disable"
}

func searchPath(value string) string {
	if value != "" {
		return fmt.Sprintf("&search_path=%s", value)
	}

	return ""
}
