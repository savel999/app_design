package repos

import (
	"context"

	"github.com/savel999/app_design/internal/domain/models"
)

type UsersRepo interface {
	GetOrCreate(ctx context.Context, in GetOrCreateQuery) (models.User, error)
	Create(ctx context.Context, in models.User) (models.User, error)
}

type GetOrCreateQuery struct {
	Email string
}
