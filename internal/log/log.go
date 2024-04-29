package log

import (
	"log/slog"
	"os"
)

var logHandler = slog.NewJSONHandler(os.Stdout, nil)

func NewLogger(service string) *slog.Logger {
	baseLogger := slog.New(logHandler)
	return baseLogger.With("zone", service)
}
