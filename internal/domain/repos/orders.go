package repos

import (
	"context"

	"github.com/savel999/app_design/internal/domain/models"
	"github.com/savel999/app_design/internal/domain/types"
)

type OrdersRepo interface {
	Create(ctx context.Context, in models.Order) (models.Order, error)
	SetStatus(ctx context.Context, id int, status types.OrderStatus) error
}
