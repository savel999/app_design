package repos

import (
	"context"

	"github.com/savel999/app_design/internal/domain/models"
)

type PaymentsRepo interface {
	Create(ctx context.Context, in models.Payment) (models.Payment, error)
}
