package logger

import (
	"context"
	"io"
	"log/slog"
	"runtime"
	"time"
)

type TraceIDFn func(ctx context.Context) string

type Logger struct {
	handler   slog.Handler
	traceIDFn TraceIDFn
}

func New(w io.Writer, minLevel Level, serviceName string, traceIDFn TraceIDFn) *Logger {
	return new(w, minLevel, serviceName, traceIDFn, Events{})
}

func NewWithEvents(w io.Writer, minLevel Level, serviceName string, traceIdFn TraceIDFn, events Events) *Logger {
	return new(w, minLevel, serviceName, traceIdFn, events)
}

func (log *Logger) Error(ctx context.Context, msg string, args ...any) {
	log.write(ctx, LevelError, 3, msg)
}

func (log *Logger) Info(ctx context.Context, msg string) {
	log.write(ctx, LevelInfo, 3, msg)
}

func (log *Logger) write(ctx context.Context, level Level, caller int, msg string) {
	slogLevel := slog.Level(level)

	var pcs [1]uintptr

	runtime.Callers(caller, pcs[:])

	r := slog.NewRecord(time.Now(), slogLevel, msg, pcs[0])
	log.handler.Handle(ctx, r)
}

func (log *Logger) Errorw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	slogLevel := slog.Level(LevelError)
	log.write(ctx, Level(slogLevel), 3, msg)
}
func new(w io.Writer, minLevel Level, serviceName string, traceIDFn TraceIDFn, events Events) *Logger {
	handler := slog.Handler(slog.NewJSONHandler(w, &slog.HandlerOptions{}))
	return &Logger{
		handler: handler,
	}
}
