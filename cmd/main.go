package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	v1 "github.com/KasperSaaby/calculatron-service/internal/api/v1"
	"github.com/KasperSaaby/calculatron-service/internal/platform/logger"
	db "github.com/KasperSaaby/calculatron-service/internal/store/database"
)

func main() {
	os.Exit(Start())
}

const (
	OK                 = 0
	FailedPrecondition = 9
	Internal           = 13
)

func Start() int {
	logger.Infof("Starting application")

	exitCode := OK
	exitChan := make(chan int, 1)

	go func() {
		exitChan <- lifetime()
	}()

	conn, err := db.New()
	if err != nil {
		logger.Errf(err, "Connect to database")
		return FailedPrecondition
	}

	err = db.MigrateSchemas(conn)
	if err != nil {
		logger.Errf(err, "Migrate database schemas")
		return FailedPrecondition
	}

	port := os.Getenv("PORT")
	if port == "" {
		logger.Errf(err, "No PORT environment variable defined")
		return FailedPrecondition
	}

	mux := http.NewServeMux()

	err = v1.Setup(mux, conn)
	if err != nil {
		logger.Errf(err, "Setup handlers")
		return FailedPrecondition
	}

	go func() {
		logger.Infof("Application listening on port %s", port)

		err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
		if err != nil {
			logger.Errf(err, "Listening on port %s", port)
		}
	}()

	logger.Infof("Application started")

	exitCode = <-exitChan

	logger.Infof("Stopping application with exit code: %d", exitCode)

	err = conn.Close()
	if err != nil {
		logger.Errf(err, "Close database")
	}

	return exitCode
}

func lifetime() int {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals

	logger.Infof("Received interrupt signal")

	return OK
}
