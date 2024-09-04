package sloger

import (
	"log/slog"
	"os"
)

var (
	logger *slog.Logger
)

func init() {
	logFile, err := os.OpenFile("logs/csv-validation.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0666)

	if err != nil {
		panic("Failed to Open the log file:" + logFile.Name())
	}

	logger = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: slog.LevelInfo}))
}

func With(args ...any) *slog.Logger {
	if logger == nil {
		return slog.With(args...)
	}
	return logger.With(args...)
}
