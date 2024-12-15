package services

import (
	"context"
	"fmt"
	"time"

	"github.com/savel999/app_design/internal/domain/models"
	"github.com/savel999/app_design/internal/domain/repos"
	"github.com/savel999/app_design/internal/domain/types"
	pkgtime "github.com/savel999/app_design/pkg/time"
)

type OrdersService interface {
	Create(ctx context.Context, q CreateOrderQuery) (models.FullOrder, error)
	Calculate(ctx context.Context, q CalculateOrderQuery) (float64, error)
}

type CalculateOrderQuery struct {
	UserID   int
	Bookings []Booking
}

type CreateOrderQuery struct {
	UserID   int
	Bookings []Booking
}

type Booking struct {
	Room models.Room
	From time.Time
	To   time.Time
}

type ordersService struct {
	bookingsRepo repos.BookingsRepo
	ordersRepo   repos.OrdersRepo
	paymentsRepo repos.PaymentsRepo
}

func NewOrdersService(
	bookingsRepo repos.BookingsRepo,
	ordersRepo repos.OrdersRepo,
	paymentsRepo repos.PaymentsRepo,
) OrdersService {
	return &ordersService{
		bookingsRepo: bookingsRepo,
		ordersRepo:   ordersRepo,
		paymentsRepo: paymentsRepo,
	}
}

func (o *ordersService) Create(ctx context.Context, in CreateOrderQuery) (models.FullOrder, error) {
	out := models.FullOrder{}
	createOrderIn := models.Order{UserID: in.UserID, CreatedAt: time.Now(), Status: types.OrderStatusNew}

	// создание заказа
	order, err := o.ordersRepo.Create(ctx, createOrderIn)
	if err != nil {
		return out, fmt.Errorf("failed to create order in storage: %w", err)
	}

	// создание броней
	for _, booking := range in.Bookings {
		createBookingIn := models.Booking{
			OrderID: order.ID,
			RoomID:  booking.Room.ID,
			From:    booking.From,
			To:      booking.To,
		}

		newBooking, err := o.bookingsRepo.Create(ctx, createBookingIn)
		if err != nil {
			return out, fmt.Errorf("failed to create order booking in storage: %w", err)
		}

		out.Bookings = append(out.Bookings, newBooking)
	}

	//рассчитать заказ
	amount, err := o.Calculate(ctx, CalculateOrderQuery{UserID: in.UserID, Bookings: in.Bookings})
	if err != nil {
		return out, fmt.Errorf("failed to calculate order amount: %w", err)
	}

	payment, err := o.paymentsRepo.Create(ctx, models.Payment{OrderID: order.ID, Amount: amount})
	if err != nil {
		return out, fmt.Errorf("failed to create order payment in storage: %w", err)
	}

	out.Payment = payment

	// заказ готов к оплате
	if err = o.ordersRepo.SetStatus(ctx, order.ID, types.OrderStatusReadyToPay); err != nil {
		return out, fmt.Errorf("failed to set order ready to pay: %w", err)
	}

	order.Status = types.OrderStatusReadyToPay
	out.Order = order

	return out, nil
}

func (o *ordersService) Calculate(_ context.Context, in CalculateOrderQuery) (float64, error) {
	amount := 0.0

	for _, booking := range in.Bookings {
		amount += booking.Room.Price * float64(pkgtime.GetDaysDifference(booking.From, booking.To))
	}

	return amount, nil
}
