package dto

import "time"

type RoomRaw struct {
	ID      int
	HotelID int
	Name    string
	Count   int
	Price   float64
}

type HotelRaw struct {
	ID       int
	Name     string
	CheckIn  time.Time
	CheckOut time.Time
}

type UserRaw struct {
	ID    int
	Email string
}

type BookingRaw struct {
	RoomID  int
	OrderID int
	From    time.Time
	To      time.Time
}

type OrderRaw struct {
	ID        int
	UserID    int
	CreatedAt time.Time
	Status    string
}

type CountRoomBookingsByPeriodQuery struct {
	RoomID int
	To     time.Time
	From   time.Time
}

type PaymentRaw struct {
	ID        int
	OrderID   int
	CreatedAt time.Time
	Amount    float64
}
