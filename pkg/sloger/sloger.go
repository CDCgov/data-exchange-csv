package sloger

import (
	"log/slog"
	"os"
)

var (
	logger *slog.Logger
)

func InitLogger(logToFile bool, debug bool) {
	/*initialize log level to Info
	if debug is true, update logLevel to Debug
	*/
	logLevel := slog.LevelInfo
	if debug {
		logLevel = slog.LevelDebug
	}
	if logToFile {
		logFile, err := os.OpenFile("logs/csv-validation.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0666)
		if err != nil {
			panic("Failed to Open the log file:" + logFile.Name())
		}
		logger = slog.New(slog.NewJSONHandler(logFile, &slog.HandlerOptions{Level: logLevel}))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))
	}
}

func With(args ...any) *slog.Logger {
	if logger == nil {
		return slog.With(args...)
	}

	return logger.With(args...)
}
