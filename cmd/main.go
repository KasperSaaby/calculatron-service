package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/KasperSaaby/calculatron-service/generated/restapi"
	"github.com/KasperSaaby/calculatron-service/generated/restapi/operations"
	v1 "github.com/KasperSaaby/calculatron-service/internal/api/v1"
	"github.com/KasperSaaby/calculatron-service/internal/app"
	"github.com/KasperSaaby/calculatron-service/internal/domain"
	"github.com/KasperSaaby/calculatron-service/internal/platform/logger"
	"github.com/KasperSaaby/calculatron-service/internal/store"
	"github.com/KasperSaaby/calculatron-service/internal/store/database"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"go.uber.org/fx"
)

func main() {
	fx.New(Opts()).Run()
}

func Opts() fx.Option {
	return fx.Options(
		fx.Provide(
			NewDatabase,
			NewSwaggerAPI,
		),
		v1.Setup,
		app.Setup,
		store.Setup,
		domain.Setup,
	)
}

func NewDatabase(lc fx.Lifecycle) (*sql.DB, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return database.MigrateSchemas(db)
		},
		OnStop: func(context.Context) error {
			return db.Close()
		},
	})

	return db, nil
}

func NewSwaggerAPI(lc fx.Lifecycle) (*operations.CalculatronServiceAPI, error) {
	spec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, fmt.Errorf("load swagger spec: %w", err)
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		return nil, fmt.Errorf("no PORT environment variable defined")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("invalid PORT environment variable %q", portStr)
	}

	api := operations.NewCalculatronServiceAPI(spec)
	api.ServeError = errors.ServeError
	api.UseSwaggerUI()

	server := restapi.NewServer(api)
	server.Port = port

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := server.Serve()
				if err != nil {
					logger.Errf(err, "Serving HTTP server")
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			return server.Shutdown()
		},
	})

	return api, nil
}
