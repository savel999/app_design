package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/savel999/app_design/internal/domain/models"
	"github.com/savel999/app_design/internal/domain/types"
	"github.com/savel999/app_design/internal/infrastructure/storage"
	"github.com/savel999/app_design/internal/infrastructure/storage/dto"
)

type OrdersRepository struct {
	ordersStore storage.Orders
}

func NewOrdersRepository(ordersStore storage.Orders) *OrdersRepository {
	return &OrdersRepository{ordersStore: ordersStore}
}

func (r *OrdersRepository) Create(ctx context.Context, q models.Order) (models.Order, error) {
	in := dto.OrderRaw{CreatedAt: time.Now(), UserID: q.UserID, Status: q.Status.String()}

	order, err := r.ordersStore.Create(ctx, in)
	if err != nil {
		return models.Order{}, fmt.Errorf("can't create order: %w", err)
	}

	return mapOrderToModel(order), nil
}

func (r *OrdersRepository) SetStatus(ctx context.Context, id int, status types.OrderStatus) error {
	if err := r.ordersStore.SetStatus(ctx, id, status.String()); err != nil {
		return fmt.Errorf("can't set order status: %w", err)
	}

	return nil
}

func mapOrderToModel(o dto.OrderRaw) models.Order {
	return models.Order{ID: o.ID, UserID: o.UserID, CreatedAt: o.CreatedAt}
}
