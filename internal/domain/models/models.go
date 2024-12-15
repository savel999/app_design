package models

import (
	"time"

	"github.com/savel999/app_design/internal/domain/types"
)

type User struct {
	ID    int
	Email string
}

type Room struct {
	ID      int
	HotelID int
	Name    string
	Count   int
	Price   float64
}

type Hotel struct {
	ID       int
	Name     string
	CheckIn  time.Time
	CheckOut time.Time
}

type Booking struct {
	RoomID  int
	OrderID int
	From    time.Time
	To      time.Time
}

type Order struct {
	ID        int
	UserID    int
	CreatedAt time.Time
	Status    types.OrderStatus
}

type Payment struct {
	ID        int
	OrderID   int
	CreatedAt time.Time
	Amount    float64
}

type FullOrder struct {
	Order    Order
	Payment  Payment
	Bookings []Booking
}
