package main

import (
	v1 "calculatron/internal/api/v1"
	"calculatron/internal/db"
	"calculatron/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	os.Exit(Start())
}

const (
	// OK indicates application exited with no issues
	OK = 0
	// FailedPrecondition occurs if setup fails
	FailedPrecondition = 9
	// Internal occurs if an error is encountered while the application is running
	Internal = 13
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

	mux := http.NewServeMux()

	err = v1.Setup(mux, conn)
	if err != nil {
		logger.Errf(err, "Setup handlers")
		return FailedPrecondition
	}

	go func() {
		err := http.ListenAndServe(":80", mux)
		if err != nil {
			logger.Errf(err, "Listening on port 80")
		}
	}()

	logger.Infof("Application started")

	exitCode = <-exitChan

	logger.Infof("Received exit code: %d", exitCode)
	logger.Infof("Stopping application")

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
