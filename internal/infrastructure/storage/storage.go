package storage

import (
	"context"

	"github.com/savel999/app_design/internal/infrastructure/storage/dto"
)

type Rooms interface {
	GetByID(ctx context.Context, id int) (dto.RoomRaw, error)
	Create(ctx context.Context, room dto.RoomRaw) (dto.RoomRaw, error)
}

type Hotels interface {
	GetByID(ctx context.Context, id int) (dto.HotelRaw, error)
	Create(ctx context.Context, room dto.HotelRaw) (dto.HotelRaw, error)
}

type Users interface {
	GetByEmail(ctx context.Context, email string) (dto.UserRaw, error)
	Create(ctx context.Context, room dto.UserRaw) (dto.UserRaw, error)
}

type Bookings interface {
	CountRoomBookingsByPeriod(ctx context.Context, q dto.CountRoomBookingsByPeriodQuery) (int, error)
	Create(ctx context.Context, room dto.BookingRaw) (dto.BookingRaw, error)
}

type Orders interface {
	Create(ctx context.Context, room dto.OrderRaw) (dto.OrderRaw, error)
	SetStatus(ctx context.Context, id int, status string) error
}

type Payments interface {
	Create(ctx context.Context, room dto.PaymentRaw) (dto.PaymentRaw, error)
}
