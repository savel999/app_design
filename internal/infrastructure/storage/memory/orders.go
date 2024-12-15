package memory

import (
	"context"
	"sync"

	"github.com/savel999/app_design/internal/infrastructure/storage"
	"github.com/savel999/app_design/internal/infrastructure/storage/dto"
)

type orders struct {
	mu     *sync.RWMutex
	data   map[int]dto.OrderRaw
	lastID int
}

func NewOrdersStorage() storage.Orders {
	return &orders{mu: &sync.RWMutex{}, data: make(map[int]dto.OrderRaw)}
}

func (s *orders) Create(_ context.Context, order dto.OrderRaw) (dto.OrderRaw, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lastID++

	order.ID = s.lastID
	s.data[s.lastID] = order

	return order, nil
}

func (s *orders) SetStatus(_ context.Context, id int, status string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, ok := s.data[id]
	if !ok {
		return storage.ErrNotFound
	}

	order.Status = status
	s.data[id] = order

	return nil
}
