package web

import (
	"fmt"
	"log/slog"

	"github.com/savel999/app_design/internal/app/web/registry"
	"github.com/savel999/app_design/internal/app/web/server"
)

func Run(config *Config) error {
	cnt := NewContainer(config)

	cnt.Logger.Debug("init config", slog.Any("config", config))

	s := InitServer(cnt)

	beforeShutdown := func() {
		cnt.Clean()
	}

	if err := s.ListenAndServe(config.ServerAddr, beforeShutdown); err != nil {
		return fmt.Errorf("failed to listen and serve: %w", err)
	}

	return nil
}

func InitServer(cnt *registry.Container) *server.Server {
	s := server.NewServer(cnt.Logger)

	s.InitMiddlewares()
	s.InitRoutes(cnt.Handlers)

	return s
}
