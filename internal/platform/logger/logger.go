package logger

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

var jsonLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func Infof(ctx context.Context, format string, args ...interface{}) {
	jsonLogger.InfoContext(ctx, format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	jsonLogger.ErrorContext(ctx, format, args...)
}

func Errf(ctx context.Context, err error, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	errMsg := "nil"
	if err != nil {
		errMsg = err.Error()
	}

	jsonLogger.ErrorContext(ctx, fmt.Sprintf("%s: %s", msg, errMsg))
}
