package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/savel999/app_design/internal/domain/models"
	"github.com/savel999/app_design/internal/infrastructure/storage"
	"github.com/savel999/app_design/internal/infrastructure/storage/dto"
)

type PaymentsRepository struct {
	paymentsStore storage.Payments
}

func NewPaymentsRepository(paymentsStore storage.Payments) *PaymentsRepository {
	return &PaymentsRepository{paymentsStore: paymentsStore}
}

func (r *PaymentsRepository) Create(ctx context.Context, q models.Payment) (models.Payment, error) {
	in := dto.PaymentRaw{CreatedAt: time.Now(), OrderID: q.OrderID, Amount: q.Amount}

	payment, err := r.paymentsStore.Create(ctx, in)
	if err != nil {
		return models.Payment{}, fmt.Errorf("can't order payment: %w", err)
	}

	return mapPaymentToModel(payment), nil
}

func mapPaymentToModel(o dto.PaymentRaw) models.Payment {
	return models.Payment{ID: o.ID, OrderID: o.OrderID, Amount: o.Amount, CreatedAt: o.CreatedAt}
}
