package logger

import (
	"context"
	"log/slog"
	"time"
)

type Level slog.Level

// A set of possible logging levels.
const (
	LevelDebug = Level(slog.LevelDebug)
	LevelInfo  = Level(slog.LevelInfo)
	LevelWarn  = Level(slog.LevelWarn)
	LevelError = Level(slog.LevelError)
)

// Record represents the data that is being logged.
type Record struct {
	Time       time.Time
	Message    string
	Level      Level
	Attributes map[string]any
}
type EventFn func(ctx context.Context, r Record)

type Events struct {
	Debug EventFn
	Info  EventFn
	Warn  EventFn
	Error EventFn
}
