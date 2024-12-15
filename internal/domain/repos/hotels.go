package repos

import (
	"context"

	"github.com/savel999/app_design/internal/domain/models"
)

type HotelsRepo interface {
	GetByID(ctx context.Context, id int) (models.Hotel, error)
	Create(ctx context.Context, in models.Hotel) (models.Hotel, error)
}
