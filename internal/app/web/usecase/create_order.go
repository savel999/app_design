package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	domainerrors "github.com/savel999/app_design/internal/domain/errors"
	"github.com/savel999/app_design/internal/domain/models"
	"github.com/savel999/app_design/internal/domain/repos"
	"github.com/savel999/app_design/internal/domain/services"
	"github.com/savel999/app_design/internal/presentation/rest/dto"
	pkgtime "github.com/savel999/app_design/pkg/time"
)

type CreateOrder func(
	ctx context.Context, input dto.CreateOrderRequest,
) (dto.CreateOrderResponse, error)

func NewCreateOrder(
	usersRepo repos.UsersRepo,
	hotelsRepo repos.HotelsRepo,
	roomsRepo repos.RoomsRepo,
	bookingsRepo repos.BookingsRepo,
	ordersService services.OrdersService,
) CreateOrder {
	return func(ctx context.Context, input dto.CreateOrderRequest) (dto.CreateOrderResponse, error) {
		out := dto.CreateOrderResponse{}

		if errs := validateCreateOrderRequest(input); len(errs) > 0 {
			return out, dto.NewValidationErrors("input validation errors", errs)
		}

		//по отелю получим время заезда-выселения
		hotel, err := hotelsRepo.GetByID(ctx, input.HotelID)
		if err != nil {
			return out, fmt.Errorf("failed to get hotel by id: %w", err)
		}

		from := pkgtime.SetClock(input.From, hotel.CheckIn.Hour(), hotel.CheckIn.Minute(), hotel.CheckIn.Second())
		to := pkgtime.SetClock(input.To, hotel.CheckOut.Hour(), hotel.CheckOut.Minute(), hotel.CheckOut.Second())

		//проверить связку отель-номер
		room, err := roomsRepo.GetByID(ctx, input.RoomID)
		switch {
		case errors.Is(err, domainerrors.ErrNotFound):
			return out, dto.NewValidationErrors("room not found", nil)
		case err == nil && room.HotelID != input.HotelID:
			return out, dto.NewValidationErrors("room not found in selected hotel", nil)
		case err != nil:
			return out, fmt.Errorf("failed to get room by id: %w", err)
		}

		//проверить свободные места
		searchBookings := repos.CountRoomBookingsByPeriodQuery{RoomID: room.ID, From: from, To: to}

		bookingsCnt, err := bookingsRepo.CountRoomBookingsByPeriod(ctx, searchBookings)
		if err != nil {
			return out, fmt.Errorf("failed to count bookings for room: %w", err)
		}

		if room.Count-bookingsCnt < 1 {
			return out, dto.NewValidationErrors("room not available for selected dates", nil)
		}

		// получить по email или создать нового пользовтеля
		user, err := usersRepo.GetOrCreate(ctx, repos.GetOrCreateQuery{Email: input.Email})
		if err != nil {
			return out, fmt.Errorf("failed to get user: %w", err)
		}

		//создать заказ
		fullOrder, err := ordersService.Create(ctx, services.CreateOrderQuery{
			UserID: user.ID, Bookings: []services.Booking{{Room: room, From: from, To: to}},
		})
		if err != nil {
			return out, fmt.Errorf("failed to create service order: %w", err)
		}

		return mapModelToCreateOrderResponseDTO(fullOrder, user), nil
	}
}

func validateCreateOrderRequest(input dto.CreateOrderRequest) []error {
	var errs []error

	if input.RoomID <= 0 {
		errs = append(errs, fmt.Errorf("room_id is required and must be > 0"))
	}

	if input.HotelID <= 0 {
		errs = append(errs, fmt.Errorf("hotel_id is required and must be > 0"))
	}

	if input.Email == "" { // упрощенная проверка
		errs = append(errs, fmt.Errorf("email is required"))
	}

	if input.From.IsZero() || input.To.IsZero() || input.From.After(input.To) || input.From.Before(time.Now()) {
		errs = append(errs, fmt.Errorf("from and to is required, from less than to, from and to more than now"))
	}

	return errs
}

func mapModelToCreateOrderResponseDTO(o models.FullOrder, u models.User) dto.CreateOrderResponse {
	return dto.CreateOrderResponse{
		ID:       o.Order.ID,
		Price:    o.Payment.Amount,
		Status:   o.Order.Status.String(),
		Email:    u.Email,
		Bookings: mapBookings(o.Bookings),
	}
}

func mapBookings(bookings []models.Booking) []dto.Booking {
	result := make([]dto.Booking, 0, len(bookings))

	for _, booking := range bookings {
		result = append(result, dto.Booking{RoomID: booking.RoomID, From: booking.From, To: booking.To})
	}

	return result
}

type CreateOrderResponse struct {
	ID     int       `json:"id"`
	RoomID int       `json:"room_id"`
	Email  string    `json:"email"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
	Price  float64   `json:"price"`
	Status string    `json:"status"`
}
