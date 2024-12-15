package memory

import (
	"context"
	"sync"

	"github.com/savel999/app_design/internal/infrastructure/storage"
	"github.com/savel999/app_design/internal/infrastructure/storage/dto"
)

type hotels struct {
	mu     *sync.RWMutex
	data   map[int]dto.HotelRaw
	lastID int
}

func NewHotelsStorage() storage.Hotels {
	return &hotels{mu: &sync.RWMutex{}, data: make(map[int]dto.HotelRaw)}
}

func (s *hotels) GetByID(_ context.Context, id int) (dto.HotelRaw, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, v := range s.data {
		if v.ID == id {
			return v, nil
		}
	}

	return dto.HotelRaw{}, storage.ErrNotFound
}

func (s *hotels) Create(_ context.Context, hotel dto.HotelRaw) (dto.HotelRaw, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID++

	hotel.ID = s.lastID
	s.data[s.lastID] = hotel

	return hotel, nil
}
