package handlers

import (
	"net/http"
	"time"
)

func NewCreateOrderHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

type CreateOrderRequest struct {
	HotelID int       `json:"hotel_id"`
	RoomID  int       `json:"room_id"`
	Email   string    `json:"email"`
	From    time.Time `json:"from"`
	To      time.Time `json:"to"`
}

type CreateOrderResponse struct {
	ID int `json:"id"`
}
