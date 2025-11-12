package main

import (
	"log/slog"
	"os"
)

// logLevel returns a [log/slog.Level] depending on value of the verbose flag.
func logLevel(verbose bool) slog.Level {
	if verbose {
		return slog.LevelDebug
	}

	return slog.LevelInfo
}

// configureLogging sets the log level and sets a default logger for the application.
func configureLogging(verbose bool) {
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: logLevel(verbose),
	})

	slog.SetDefault(slog.New(h))
}
