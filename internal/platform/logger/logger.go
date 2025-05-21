package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	stdoutLogger = log.New(os.Stdout, "log: ", log.LstdFlags)
	stderrLogger = log.New(os.Stderr, "log: ", log.LstdFlags)
)

func Infof(format string, args ...interface{}) {
	stdoutLogger.Printf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	stderrLogger.Printf(format, args...)
}

func Errf(err error, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	errMsg := "nil"
	if err != nil {
		errMsg = err.Error()
	}

	stderrLogger.Printf("%s: %s", msg, errMsg)
}
