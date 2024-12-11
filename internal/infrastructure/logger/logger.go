package logger

import (
	"context"
	"log/slog"
)

type Logger interface {
	InfoContext(ctx context.Context, msg string, fields ...slog.Attr)
	Info(msg string, fields ...slog.Attr)

	ErrorContext(ctx context.Context, msg string, fields ...slog.Attr)
	Error(msg string, fields ...slog.Attr)
}

func New() Logger {

}
