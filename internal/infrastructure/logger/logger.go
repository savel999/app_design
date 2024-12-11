package logger

import (
	"context"
	"log/slog"
	"os"
)

const DefaultAttrGroupName = "log"

type Logger interface {
	DebugContext(ctx context.Context, msg string, attrs ...slog.Attr)
	Debug(msg string, attrs ...slog.Attr)

	InfoContext(ctx context.Context, msg string, attrs ...slog.Attr)
	Info(msg string, attrs ...slog.Attr)

	WarnContext(ctx context.Context, msg string, attrs ...slog.Attr)
	Warn(msg string, attrs ...slog.Attr)

	ErrorContext(ctx context.Context, msg string, attrs ...slog.Attr)
	Error(msg string, attrs ...slog.Attr)
}

func New(LogLevel slog.Level) Logger {
	return &logger{
		logger: slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: LogLevel}),
		),
	}
}

type logger struct {
	logger *slog.Logger
}

func (l *logger) DebugContext(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.logAttrs(ctx, slog.LevelDebug, msg, attrs...)
}

func (l *logger) Debug(msg string, attrs ...slog.Attr) {
	l.logAttrs(context.Background(), slog.LevelDebug, msg, attrs...)
}

func (l *logger) InfoContext(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.logAttrs(ctx, slog.LevelInfo, msg, attrs...)
}

func (l *logger) Info(msg string, attrs ...slog.Attr) {
	l.logAttrs(context.Background(), slog.LevelInfo, msg, attrs...)
}

func (l *logger) WarnContext(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.logAttrs(ctx, slog.LevelWarn, msg, attrs...)
}

func (l *logger) Warn(msg string, attrs ...slog.Attr) {
	l.logAttrs(context.Background(), slog.LevelWarn, msg, attrs...)
}

func (l *logger) ErrorContext(ctx context.Context, msg string, attrs ...slog.Attr) {
	l.logAttrs(ctx, slog.LevelError, msg, attrs...)
}

func (l *logger) Error(msg string, attrs ...slog.Attr) {
	l.logAttrs(context.Background(), slog.LevelError, msg, attrs...)
}

func (l *logger) logAttrs(ctx context.Context, level slog.Level, msg string, attrs ...slog.Attr) {
	l.logger.LogAttrs(ctx, level, msg, slog.Attr{Key: DefaultAttrGroupName, Value: slog.GroupValue(attrs...)})
}

func ErrorAttr(err error) slog.Attr {
	return slog.Attr{Key: "error", Value: slog.AnyValue(err.Error())}
}
