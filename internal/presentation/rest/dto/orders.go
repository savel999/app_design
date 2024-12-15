package dto

import "time"

type CreateOrderRequest struct {
	HotelID int       `json:"hotel_id"`
	RoomID  int       `json:"room_id"`
	Email   string    `json:"email"`
	From    time.Time `json:"from"`
	To      time.Time `json:"to"`
}

type CreateOrderResponse struct {
	ID       int       `json:"id"`
	Price    float64   `json:"price"`
	Status   string    `json:"status"`
	Email    string    `json:"email"`
	Bookings []Booking `json:"bookings"`
}

type Booking struct {
	RoomID int       `json:"room_id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}
