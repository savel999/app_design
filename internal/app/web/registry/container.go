package registry

import (
	"github.com/savel999/app_design/internal/infrastructure/logger"
	"github.com/savel999/app_design/internal/presentation/rest"
)

type Container struct {
	Logger   logger.Logger
	Handlers *rest.Handler
}

func (cnt *Container) Clean() {
	//clean resources
}
