package memory

import (
	"context"
	"sync"

	"github.com/savel999/app_design/internal/infrastructure/storage"
	"github.com/savel999/app_design/internal/infrastructure/storage/dto"
)

type bookings struct {
	mu   *sync.RWMutex
	data []dto.BookingRaw
}

func NewBookingsStorage() storage.Bookings {
	return &bookings{mu: &sync.RWMutex{}}
}

func (s *bookings) CountRoomBookingsByPeriod(_ context.Context, q dto.CountRoomBookingsByPeriodQuery) (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	count := 0

	for _, v := range s.data {
		if v.RoomID == q.RoomID {
			if q.From.Compare(v.From) <= 0 && q.To.Compare(v.From) >= 0 ||
				q.From.Compare(v.To) <= 0 && q.To.Compare(v.To) >= 0 {
				count++
			}
		}
	}

	return count, nil
}

func (s *bookings) Create(_ context.Context, booking dto.BookingRaw) (dto.BookingRaw, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = append(s.data, booking)

	return booking, nil
}
