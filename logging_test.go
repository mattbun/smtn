package main

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogLevel_Default(t *testing.T) {
	assert.Equal(t, slog.LevelInfo, logLevel(false))
}

func TestLogLevel_Verbose(t *testing.T) {
	assert.Equal(t, slog.LevelDebug, logLevel(true))
}
