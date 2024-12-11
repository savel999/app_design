package web

import (
	"log/slog"
	"os"
	"strings"

	"github.com/savel999/app_design/pkg/env"
)

type Config struct {
	ServiceName      string
	LogLevel         slog.Level
	ServerAddr       string
	WithPreparedData bool
}

func InitConfig() *Config {
	return &Config{
		ServiceName:      env.GetString("SERVICE_NAME", "web"),
		ServerAddr:       env.GetString("SERVER_ADDR", ":8080"),
		LogLevel:         GetLogLevel("LOG_LEVEL", slog.LevelInfo),
		WithPreparedData: env.GetBool("WITH_PREPARED_DATA", true),
	}
}

func GetLogLevel(name string, defaultVal slog.Level) slog.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(name))) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return defaultVal
	}
}
