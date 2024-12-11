package web

import (
	"github.com/savel999/app_design/internal/app/web/registry"
	"github.com/savel999/app_design/internal/infrastructure/logger"
	"github.com/savel999/app_design/internal/presentation/rest"
)

func NewContainer(
	config *Config,
) *registry.Container {
	log := logger.New(config.LogLevel)

	return &registry.Container{
		Logger:   log,
		Handlers: rest.New(log),
	}
}
