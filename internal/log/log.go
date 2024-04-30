package log

import (
	"log/slog"
	"os"
	"strings"
)

var logHandler = slog.NewTextHandler(os.Stdout, nil)

func NewLogger(service string) *slog.Logger {
	baseLogger := slog.New(logHandler)
	return baseLogger.With("zone", strings.ToUpper(service))
}
