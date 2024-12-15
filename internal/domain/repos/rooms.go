package repos

import (
	"context"

	"github.com/savel999/app_design/internal/domain/models"
)

type RoomsRepo interface {
	GetByID(ctx context.Context, id int) (models.Room, error)
	Create(ctx context.Context, in models.Room) (models.Room, error)
}
