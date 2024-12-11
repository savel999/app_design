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
	ID     int       `json:"id"`
	RoomID int       `json:"room_id"`
	Email  string    `json:"email"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
	Price  float64   `json:"price"`
	Status string    `json:"status"`
}
