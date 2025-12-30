package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func InitLogger() error {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	Logger = slog.New(handler)
	return nil
}
