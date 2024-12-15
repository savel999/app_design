package repos

import (
	"context"
	"time"

	"github.com/savel999/app_design/internal/domain/models"
)

type BookingsRepo interface {
	Create(ctx context.Context, in models.Booking) (models.Booking, error)
	CountRoomBookingsByPeriod(ctx context.Context, q CountRoomBookingsByPeriodQuery) (int, error)
}

type CountRoomBookingsByPeriodQuery struct {
	RoomID int
	From   time.Time
	To     time.Time
}
