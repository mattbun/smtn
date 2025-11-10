package main

import (
	"log/slog"
	"os"
)

func logLevel(verbose bool) slog.Level {
	if verbose {
		return slog.LevelDebug
	}

	return slog.LevelInfo
}

func configureLogging(verbose bool) {
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel(verbose),
	})

	slog.SetDefault(slog.New(h))
}
