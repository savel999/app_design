package storage

import (
	"context"
	"fmt"

	"github.com/savel999/app_design/internal/domain/models"
	"github.com/savel999/app_design/internal/domain/repos"
	"github.com/savel999/app_design/internal/infrastructure/storage"
	"github.com/savel999/app_design/internal/infrastructure/storage/dto"
)

type BookingsRepository struct {
	bookingsStore storage.Bookings
}

func NewBookingsRepository(bookingsStore storage.Bookings) *BookingsRepository {
	return &BookingsRepository{bookingsStore: bookingsStore}
}

func (r *BookingsRepository) Create(ctx context.Context, q models.Booking) (models.Booking, error) {
	in := dto.BookingRaw{
		RoomID:  q.RoomID,
		OrderID: q.OrderID,
		From:    q.From,
		To:      q.To,
	}

	booking, err := r.bookingsStore.Create(ctx, in)
	if err != nil {
		return models.Booking{}, fmt.Errorf("can't create booking: %w", err)
	}

	return mapBookingToModel(booking), nil
}

func (r *BookingsRepository) CountRoomBookingsByPeriod(
	ctx context.Context, q repos.CountRoomBookingsByPeriodQuery,
) (int, error) {
	in := dto.CountRoomBookingsByPeriodQuery{RoomID: q.RoomID, From: q.From, To: q.To}

	cnt, err := r.bookingsStore.CountRoomBookingsByPeriod(ctx, in)
	if err != nil {
		return 0, fmt.Errorf("can't count room bookings: %w", err)
	}

	return cnt, nil
}

func mapBookingToModel(m dto.BookingRaw) models.Booking {
	return models.Booking{RoomID: m.RoomID, OrderID: m.OrderID, From: m.From, To: m.To}
}
