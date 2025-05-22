package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"

	v1 "github.com/KasperSaaby/calculatron-service/internal/api/v1"
	"github.com/KasperSaaby/calculatron-service/internal/platform/logger"
	"github.com/KasperSaaby/calculatron-service/internal/store/database"
	"go.uber.org/fx"
)

func main() {
	fx.New(Opts()).Run()
}

func Opts() fx.Option {
	return fx.Options(
		fx.Provide(
			NewDatabase,
			NewHTTPServer,
		),
		fx.Invoke(
			database.MigrateSchemas,
			v1.Setup,
		),
	)
}

func NewDatabase(lc fx.Lifecycle) (*sql.DB, error) {
	db, err := database.New()
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return db.Close()
		},
	})

	return db, nil
}

func NewHTTPServer(lc fx.Lifecycle) *http.ServeMux {
	mux := http.NewServeMux()

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			port := os.Getenv("PORT")
			if port == "" {
				return fmt.Errorf("no PORT environment variable defined")
			}

			go func() {
				logger.Infof("Application listening on port %s", port)
				err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
				if err != nil {
					logger.Errf(err, "Listening on port %s", port)
				}
			}()

			return nil
		},
	})

	return mux
}
