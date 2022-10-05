package gorad

import (
	"context"

	"go.uber.org/zap"
)

type Logger = zap.Logger

var ctxLogger struct{}

func NewLoggerContext(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, ctxLogger, logger)
}

func LoggerFromContext(ctx context.Context) *Logger {
	return ctx.Value(ctxLogger).(*zap.Logger)
}

func NamedLoggerFromContext(ctx context.Context, name string) *Logger {
	return LoggerFromContext(ctx).Named(name)
}

func NewNamedLoggerFromContext(ctx context.Context, name string) context.Context {
	return NewLoggerContext(ctx, NamedLoggerFromContext(ctx, name))
}
